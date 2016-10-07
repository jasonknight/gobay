<?php
function xmlObjToArr($obj) { 
    $namespace = $obj->getDocNamespaces(true); 
    $namespace[NULL] = NULL; 
    
    $children = array(); 
    $attributes = array(); 
    $name = (string)$obj->getName(); 
    
    $text = trim((string)$obj); 
    if( strlen($text) <= 0 ) { 
        $text = NULL; 
    } 
    
    // get info for all namespaces 
    if(is_object($obj)) { 
        foreach( $namespace as $ns=>$nsUrl ) { 
            // atributes 
            $objAttributes = $obj->attributes($ns, true); 
            foreach( $objAttributes as $attributeName => $attributeValue ) { 
                $attribName = \trim((string)$attributeName); 
                $attribVal = trim((string)$attributeValue); 
                if (!empty($ns)) { 
                    $attribName = $ns . ':' . $attribName; 
                } 
                $attributes[$attribName] = $attribVal; 
            } 
            
            // children 
            $objChildren = $obj->children($ns, true); 
            foreach( $objChildren as $childName=>$child ) { 
                $childName = (string)$childName; 
                if( !empty($ns) ) { 
                    $childName = $ns.':'.$childName; 
                } 
                $children[$childName][] = xmlObjToArr($child); 
            } 
        } 
    } 
    
    return array( 
        'name'=>$name, 
        'text'=>$text, 
        'attributes'=>$attributes, 
        'children'=>$children 
    ); 
} 

$example = $_SERVER['argv'][2]; // the example XML File
$pkg = $_SERVER['argv'][1];
$parts = explode(".", basename($example));

if ( !file_exists($example) ) {
    die("$example does not exist!\n");
} 
libxml_use_internal_errors(true);
$example_contents = file_get_contents($example);
$r = simplexml_load_string($example_contents);

if ($r === false) {
    echo "Failed loading XML\n";
    foreach(libxml_get_errors() as $error) {
        echo "\t", $error->message;
    }
}

$r = xmlObjToArr($r);
$structs = array();
function cleanName($n) {
    $n = str_replace("Get","",$n);
    // Response -> Result
    $n = str_replace("Response","Result",$n);
    return $n;
}
function toSnakeCase($in) {
  preg_match_all('!([A-Z][A-Z0-9]*(?=$|[A-Z][a-z0-9])|[A-Za-z][a-z0-9]+)!', $in, $matches);
  $ret = $matches[0];
  foreach ($ret as &$match) {
    $match = $match == strtoupper($match) ? strtolower($match) : lcfirst($match);
  }
  return implode('_', $ret);
}
function discernType($xtxt,$v=array()) {
    $type = "string";
    if (is_numeric($xtxt)) {
        if (is_double($xtxt)) {
            $type = "float32";
        } 
        $type = "int64";
    } 
    
    if ( count($v) > 1) {
        $type = "[]$type";
    }
    return $type;
}
// The root will be the call name
function to_struct($r,$name="") {
    global $structs;
    if (empty($name)) {
        $struct = $r['name'];
    } else {
        $struct = $name;
    }
    // we remove the Get part
    $struct = cleanName($struct);
    $structs[$struct] = array(
        'def' => '',
        'attributes' => array(),
    );
    // now we need to collect up the children
    $txt = "type $struct struct{\n";
    $nodes_with_children = array();
    foreach ($r['children'] as $k=>$v) {
        $cname = $k;
        $scname = $k;
        $type = discernType($v[0]['text'],$v);
        if(count($v[0]['children']) > 0 && count($v[0]['attributes']) == 0) {
            if ( $cname[strlen($cname)-1] == 's') {
               $type = substr($cname,0,strlen($cname)-1); 
            } elseif (preg_match("/Array/",$cname)) {
               $type = str_replace("Array","",$cname);
            } else {
                $type = $cname;
            }
            $nodes_with_children[$type] = $v;
            if ( count($v) > 1) {
                $type = "[]$type";
            }
        }
        if ( count($v[0]['attributes']) > 0) {
            $otype = $type;
            $type = "struct {\n";
            foreach ($v[0]['attributes'] as $a=>$val) {

                $aname = ucfirst($a);
                $type .= "\t\t$aname " . discernType($val) . " `xml:\"$a,attr\" yaml:\"$a\"`\n";
                $type .= "\t\tValue $otype `xml:\",chardata\" yaml:\"Value\"`\n";
            }
            $type .= "\t}";
            $txt .= "\t$cname $type `xml:\"$cname\" yaml:\"$cname\"`\n";
        } else {
            $txt .= "\t$cname $type `xml:\"$cname\" yaml:\"$cname\"`\n";
        }
        //echo "\t$cname $type `xml:\"$cname\"`\n";
        
        $structs[$struct]['attributes'][$cname] = array($type,$v);
    } 
    $txt .= "}\n";
    $structs[$struct]["def"] = $txt;
    foreach ($nodes_with_children as $n=>$ch) {
        to_struct($ch[0],$n);
    }
}
to_struct($r);

// now we need to build the def file

$target = toSnakeCase(cleanName($r['name']));
echo "Target file is: $target\n";
$target_txt = "package $pkg\n";
$target_txt .= "import (\n\t\"encoding/xml\"\n";
$target_txt .= "\t\"gopkg.in/yaml.v2\"\n";
$target_txt .= ")\n";
$i = 0;
foreach ( $structs as $name => $desc ) {
    $target_txt .= $desc['def'] . "\n";
    // Now we create the new function:
    if ($i == 0) { // only for the root!
        $nf = "func New{$name}() *{$name} {\n";
        $nf .= "\treturn &{$name}{}\n";
        $nf .= "}\n";
        $target_txt .= $nf . "\n";
        // Now we create a FromXML function
        $fxml = "func (o *{$name}) FromXML(data []byte) error {\n";
        $fxml .= "\t return xml.Unmarshal(data,o)\n";
        $fxml .= "}\n";
        $target_txt .= $fxml . "\n";
        // Now we create a FromYAML function
        $fxml = "func (o *{$name}) FromYAML(data []byte) error {\n";
        $fxml .= "\t return yaml.Unmarshal(data,o)\n";
        $fxml .= "}\n";
        $target_txt .= $fxml . "\n";
    }
    foreach ($desc['attributes'] as $a=>$t) {
        if ($a == 'Ack') {
            // this is a result so need general success/failure
            $funcs = array("Success","Failure","Warning");
            foreach ($funcs as $func) {
                $bf = "func (o *{$name}) {$func}() bool {\n";
                $bf .= "\tif o.{$a} == \"$func\" {\n";
                $bf .= "\t\treturn true\n";
                $bf .= "\t}\n";
                $bf .= "\treturn false\n";
                $bf .= "}\n";
                $target_txt .= $bf . "\n";
            }
        }
    }
    $i++;
}

//echo $target_txt;

file_put_contents($target . ".go", $target_txt);

// Now we need to make an automated test



function base_type($t) {
    return str_replace("[]","",$t);
}
function gen_test_funcs($name,$desc,$receiver="") {
    global $txt;
    global $structs;
    $i = 0;
    if ( base_type($name) != $name) {
        $receiver = substr($receiver,0,strlen($receiver) - 1) . "[$i].";
    }
    foreach ($desc['attributes'] as $a=>$t) {
        $tval = $t[1][0]['text'];
        $values = $t[1];
        $bf = "";
        
        if ( $t[0] == 'string' ) {
            $bf .= "\tif o.{$receiver}{$a} != \"{$tval}\" {\n";
            $bf .= "\t\tt.Errorf(\"failed because o.{$receiver}{$a} != {$tval}\")\n";
            $bf .= "\t\treturn\n";
            $bf .= "\t}\n"; 
        } elseif ( $t[0] == 'int64' || $t[0] == 'float32' ) {
            $bf = "\tif o.{$receiver}{$a} != {$tval} {\n";
            $bf .= "\t\tt.Errorf(\"failed because o.{$receiver}{$a} != {$tval}\")\n";
            $bf .= "\t\treturn\n";
            $bf .= "\t}\n"; 
        } else if (isset($structs[base_type($t[0])])) {
            gen_test_funcs($t[0],$structs[base_type($t[0])],"{$receiver}{$a}.");
        } else {
            continue;
        }
        
        $txt .= $bf . "\n";
    }
   
}

// print_r($structs);
// die();
$txt = "package $pkg\n";
$txt .= "import \"testing\"\n";
foreach ( $structs as $name => $desc ) {
    $txt .= "func Test{$name} ( t *testing.T ) {\n";
    $txt .= "\tdata := `$example_contents`\n";
    $txt .= "\tvar o $name\n";
    $txt .= "\terr := o.FromXML([]byte(data))\n";
    $txt .= "\tif err != nil {\n";
    $txt .= "\t\tt.Errorf(\"failed loading xml %v\",err)\n";
    $txt .= "\t\treturn\n";
    $txt .= "\t}\n";
    gen_test_funcs($name,$desc);
    $txt .= "}\n";
    break;
}
file_put_contents($target . "_test.go", $txt);