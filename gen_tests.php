<?php

function maybe_create_test_body($funcs,$desc) {

    return "";
}
$file = $_SERVER['argv'][1];
$pkg = "gobay";
// figure out the test file
    $parts = explode(".",basename($file));
    $test_file = $parts[0] . "_test.go";

$lines = file($file);
$pattern = '/^func (?P<name>[A-Za-z_0-9]+)\s*\\((?P<arglist>.+)\\)\s+(?P<return>.+)\s+\\{$/';
$pattern2 = '/^func \\((?P<recv>[A-Za-z_0-9\\*\\s]+)\\) (?P<name>[A-Za-z_0-9]+)\s*\\((?P<arglist>.+)\\)\s+(?P<return>.+)\s+\\{$/';
$pattern3 = '/^func \\((?P<recv>[A-Za-z_0-9\\*\\s]+)\\) (?P<name>[A-Za-z_0-9]+)\s*\\((?P<arglist>.+)\\)\s+\\{$/';
$funcs = array();
foreach ( $lines as $line ) {
    preg_match($pattern, $line, $m);
    if ( empty($m) ) {
        preg_match($pattern2, $line, $m);
        if ( empty($m) ) {
            preg_match($pattern3, $line, $m);
            if ( empty($m) ) {
                continue;
            }
        }
    }
    $funcs[] = $m;
}
if ( ! file_exists($test_file) ) {
    echo "Testing file doesn't exist\n";
    file_put_contents($test_file, "package $pkg\nimport \"testing\"\n");
    goto add_functions;
} else {
    echo "$test_file exists!\n";
}
print_r($funcs);
$lines = file($test_file);
$bpat = "/^func Test/";
foreach ( $lines as $line ) {
    preg_match($bpat, $line,$m);
    if ( empty($m) ) {
        continue;
    }
    foreach ($funcs as &$d ) {
        $name = "Test" . $d['name'];
        $pat = "/^func $name/";
        preg_match($pat,$line,$m);
        if ( empty( $m ) ) {
            continue;
        } 
        $d['found'] = true;
    }
}

add_functions:
$to_append = "";
foreach ( $funcs as $desc ) {
    if ( isset($desc['found']) && $desc['found'] == true ) {
        continue;
    }
    $body = maybe_create_test_body($funcs,$desc);
    $name = "Test" . $desc['name'];

    $t = "func $name(t *testing.T) {\n";
    if ( empty($body) ) {
        $t .= "\tt.Errorf(\"$name has not been filled out!!\\n\")\n";
    } else {
        $t .= $body;
    }
    
    $t .= "}\n";

    $to_append .= "$t\n";
}
//file_put_contents($test_file,$to_append,FILE_APPEND);