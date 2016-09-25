<?php

$lines = file("gobay.go");
$pattern = '/^type (?P<obj>[A-Za-z_0-9]+) struct \\{$/';
$conpat = '/^type (?P<name>[A-Za-z_0-9]+)\s+(?P<type>[A-Za-z_0-9]+)\s*$/';
$debugs = array();
$in_struct = false;
$cobject = "";
$pkg = "gobay";
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
$the_funcs = array();
foreach ( $debugs as $struct => $val ) {
    $name = "o_{$struct}";
    $attrs = $val['attrs'];
    // first we need a new func

    $func = "func New$struct () *$struct {\n";
    $func .= "\treturn &$struct{}\n";
    $func .= "}\n";
    $the_funcs[] = $func;

    // Now we need a clone function 

    $func = "func (o *$struct) Clone () *$struct {\n";
    $func .= "\tvar no $struct\n";
    foreach ($attrs as $attr => $t ) {
        $func .= "\tno.$attr = o.$attr\n";
    } 
    $func .= "\treturn &no\n";
    $func .= "}\n";
    $the_funcs[] = $func;

    // now the getters and setters

    foreach ($attrs as $attr => $t ) {
        // now let's see if it's a custom type
        $apat = "/^\\[\\](?P<ctype>[A-Za-z_0-9]+)/";
        preg_match($apat, $t,$m);
        if ( ! empty( $m ) ) {
            $ft = $m['ctype'];
            // this is an array, so it needs add and remove
            $func = "func (o *$struct) Add$ft(v $ft) {\n";
            $func .= "\to.$attr = append(o.$attr,v)\n";
            $func .= "}\n";
            $the_funcs[] = $func;

            $func = "func (o *$struct) Remove$ft(i int) {\n";
            $func .= "\tif i > len(o.$attr) {\n";
            $func .= "\t\tpanic(fmt.Sprintf(\"i:%d is out of bounds for %s.%s(%d)!\\n\"),\"$struct\",\"$attr\",len(o.$attr))\n";
            $func .= "\t}\n";
            $func .= "\to.$attr = o.{$attr}[:i+copy(o.{$attr}[i:], o.{$attr}[i+1:])]\n";
            $func .= "}\n";
            $the_funcs[] = $func;

            $func = "func (o *$struct) Get$ft(i int) $ft {\n";
            $func .= "\tif i > len(o.$attr) {\n";
            $func .= "\t\tpanic(fmt.Sprintf(\"i:%d is out of bounds for %s.%s(%d)!\\n\"),\"$struct\",\"$attr\",len(o.$attr))\n";
            $func .= "\t}\n";
            $func .= "\treturn o.{$attr}[i]\n";
            $func .= "}\n";
            $the_funcs[] = $func;
        }
        $func = "func (o *$struct) Set$attr (v $t) {\n";
        $func .= "\to.$attr = v\n";
        $func .= "}\n";
        $the_funcs[] = $func;

        $func = "func (o *$struct) Get$attr () $t {\n";
        $func .= "\treturn o.$attr\n";
        $func .= "}\n";
        $the_funcs[] = $func;
    } 



    // $func = "func ($name *{$struct}) Debug() string {\n\tvar txt string\n";
    // $attrs = $val['attrs'];
    // foreach ($attrs as $attr => $t ) {
    //     if ( $t == 'string' ) {
    //         $func .= "\ttxt = fmt.Sprintf(\"%s$struct.$attr: %s\\n\",txt, $name.$attr)\n";
    //     }
    //     // now let's see if it's a custom type
    //     $apat = "/^\\[\\](?P<ctype>[A-Za-z_0-9]+)/";
    //     preg_match($apat, $t,$m);
    //     if ( ! empty( $m ) ) {
    //         $ft = $m['ctype'];
    //         $func .= "\tfor _,v := range $name.$attr {\n";
    //         if ( isset($debugs[$ft]) ) {
    //             $func .= "\t\ttxt = fmt.Sprintf(\"%s%s\\n\",txt, v.Debug())\n";
    //         } else if ( $ft == 'string' ) {
    //             $func .= "\t\ttxt = fmt.Sprintf(\"%s%s\\n\",txt, v)\n";
    //         }

    //         $func .= "\t}\n";
    //     }
    //     // now a map
    //     $apat = "/^map\\[(?P<ktype>[A-Za-z_0-9]+)\\](?P<vtype>[A-Za-z_0-9]+)/";
    //     preg_match($apat, $t,$m);
    //     if ( ! empty($m) ) {
    //         $kt = $m['ktype'];
    //         $vt = $m['vtype'];
    //         if ( $kt == 'string' ) {
    //             $kt = "%s";
    //         }
    //         if ( $vt == 'string' ) {
    //             $vt = "%s";
    //         }
    //         $func .= "\tfor k,v := range $name.$attr {\n";
    //         $func .= "\t\ttxt = fmt.Sprintf(\"%s$struct.$attr [$kt]: $vt\\n\",txt,k,v)\n";
    //         $func .= "\t}\n";
    //     }
    // } 
    // $func .= "\treturn txt\n}\n";
}
echo "package $pkg\nimport \"fmt\"\n";
echo join("\n",$the_funcs);