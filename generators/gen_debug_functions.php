<?php
$fname = $_SERVER['argv'][2];
$lines = file($fname);
$pattern = '/^type (?P<obj>[A-Za-z_0-9]+) struct \\{$/';
$conpat = '/^type (?P<name>[A-Za-z_0-9]+)\s+(?P<type>[A-Za-z_0-9]+)\s*$/';
$debugs = array();
$in_struct = false;
$cobject = "";
$pkg = $_SERVER['argv'][1];
$tconversions = array();
foreach ( $lines as $line ) {
    preg_match($pattern, $line, $m);
    if ( ! empty( $m ) ) {
        $cobject = $m['obj'];
        $debugs[$cobject] = array();
        $in_struct = true;
    }
    preg_match($conpat, $line, $m);
    if ( ! empty( $m ) ) {
        $tname = $m['name'];
        $ttype = $m['type'];
        $tconversions[$tname] = $ttype;
    }
    if ( $in_struct == true ) {
        $apat = '/^\s*(?P<attr>[A-Za-z_0-9]+)\s+(?P<type>[\\[\\]A-Za-z_0-9]+)\s*$/';
        preg_match($apat,$line,$m);
        if ( ! empty($m) ) {
            if ( !isset($debugs[$cobject]['attrs'])) {
                $debugs[$cobject]['attrs'] = array();
            }
            $attr = $m['attr'];
            $t = $m['type'];
            $debugs[$cobject]['attrs'][$attr] = $t;
        }
    }

    preg_match('/^}\s*$/', $line, $m);
    if ( $in_struct == true && ! empty($m) ) {
        //echo "Found the end of $cobject\n";
        $in_struct = false;
    }
}
//print_r($tconversions);
$the_funcs = array();
foreach ( $debugs as $struct => $val ) {
    $name = "o_{$struct}";
    $func = "func ($name *{$struct}) Debug() string {\n\tvar txt string\n";
    $attrs = $val['attrs'];
    foreach ($attrs as $attr => $t ) {
        if ( $t == 'string' ) {
            $func .= "\ttxt = fmt.Sprintf(\"%s$struct.$attr: %s\\n\",txt, $name.$attr)\n";
        }
        // now let's see if it's a custom type
        $apat = "/^\\[\\](?P<ctype>[A-Za-z_0-9]+)/";
        preg_match($apat, $t,$m);
        if ( ! empty( $m ) ) {
            $ft = $m['ctype'];
            $func .= "\tfor _,v := range $name.$attr {\n";
            if ( isset($debugs[$ft]) ) {
                $func .= "\t\ttxt = fmt.Sprintf(\"%s%s\\n\",txt, v.Debug())\n";
            } else if ( 
                $ft == 'string' || 
                ( isset($tconversions[$ft]) && $tconversions[$ft] == 'string') 
            ) {
                $func .= "\t\ttxt = fmt.Sprintf(\"%s%s\\n\",txt, v)\n";
            }

            $func .= "\t}\n";
        }
        // now a map
        $apat = "/^map\\[(?P<ktype>[A-Za-z_0-9]+)\\](?P<vtype>[A-Za-z_0-9]+)/";
        preg_match($apat, $t,$m);
        if ( ! empty($m) ) {
            $kt = $m['ktype'];
            $vt = $m['vtype'];
            if ( $kt == 'string' ) {
                $kt = "%s";
            }
            if ( $vt == 'string' ) {
                $vt = "%s";
            }
            $func .= "\tfor k,v := range $name.$attr {\n";
            $func .= "\t\ttxt = fmt.Sprintf(\"%s$struct.$attr [$kt]: $vt\\n\",txt,k,v)\n";
            $func .= "\t}\n";
        }
    } 
    $func .= "\tglobalDebugFunction(DBG_DEBUG, txt)\n";
    $func .= "\treturn txt\n}\n";
    $the_funcs[] = $func;
}
echo "package $pkg\nimport \"fmt\"\n";
echo join("\n",$the_funcs);