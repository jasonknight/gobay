<?php
$target = $_SERVER['argv'][1]; // where to put the code
$example = $_SERVER['argv'][2]; // the xml sample
$pkg = "gobay";
$parts = explode(".", basename($example));

$struct_name = ucfirst($target);

if ( !file_exists($example) ) {
    die("$example does not exist!\n");
} 

$example_contents = file_get_contents($example);
$r = simplexml_load_string($example_contents);



function to_struct($r, $sname, $ws="") {
    $txt = "{$ws}$sname struct {\n";
    $seen = array();
    foreach ($r as $k=>$v) {
        if ( count($v) == 0 ) {
            // this is a single value
            $v = "$v";
            if ( is_numeric($v) ) {
                $t = "int64";
            } else {
                $t = "string";
            }
            $txt .= "$ws\t$k $t\n";
        } else {
            // this is an array
            $s = to_struct($v,$k,"$ws\t");
            if (isset($seen[$k]) && strlen($seen[$k] <= strlen($s)) ) {
                continue;
            }
            $seen[$k] = $s;
        }
    }
    foreach ($seen as $k=>$s) {
        $txt .= "\t{$ws}$k []" . ltrim(str_replace($k,'',$s)) ."\n";
    }
    $txt .= $ws . "}\n";
    return $txt;
}

function create_test($xml,$r,$name) {
    $txt = "func Test{$name}(t *testing.T) {\n";
    $txt .= "\tdata := `$xml`\n";
    $txt .= "\to := New{$name}([]byte(data))\n";
    $ws = "";
    $i = 0;
    foreach ($r as $k=>$v) {
        if ( count($v) == 0 ) {
            // this is a single value
            $v = "$v";
            if ( is_numeric($v) ) {
                $t = "int64";
            } else {
                $t = "string";
            }
            //$txt .= "$ws\t$k $t\n";
            if ( $t == 'string' ) {
                $tv = "\"$v\"";
            } else {
                $tv = $v;
            }
            $txt .= "\tif o.$k != $tv {\n";
            $txt .= "\t\tt.Errorf(\"$name has not been filled out %+v!!\\n\",o.$k)\n";
            $txt .= "\t}\n";

        } else {
            // this is an array
            $txt .= create_sub_test($name, "o.{$k}[$i]",$v);
            $i++;
        }
    }
    $txt .= "}\n";
    return $txt;
}
function create_sub_test($name,$recv, $r) {
    $txt = "";
   
    foreach ($r as $k=>$v) {
        if ( count($v) == 0 ) {
            // this is a single value
            $v = "$v";
            if ( is_numeric($v) ) {
                $t = "int64";
            } else {
                $t = "string";
            }
            //$txt .= "$ws\t$k $t\n";
            if ( $t == 'string' ) {
                $tv = "\"$v\"";
            } else {
                $tv = $v;
            }
            $txt .= "\tif {$recv}.{$k} != $tv {\n";
            $txt .= "\t\tt.Errorf(\"{$recv}.{$k} has not been filled out!! %+v\\n\",$recv)\n";
            $txt .= "\t}\n";
        } else {
            // this is an array
            //$txt .= to_struct($v,$k,"$ws\t");
        }
        
    }
    return $txt;
}

print_r($r->Errors);
ob_start();
echo "package $pkg\n";
echo "import \"encoding/xml\"\n";
echo "type " . to_struct($r,$struct_name,"");

echo <<<EOT
func New{$struct_name}(data []byte) *$struct_name {
    var o $struct_name
    xml.Unmarshal(data, &o)
    return &o
}

EOT;

$contents = ob_get_contents();

file_put_contents($target . ".go", $contents);
ob_end_clean();

ob_start();
echo "package $pkg\n";
echo "import \"testing\"\n";


echo create_test($example_contents,$r,$struct_name);

$contents = ob_get_contents();

file_put_contents($target . "_test.go", $contents);
ob_end_clean();

