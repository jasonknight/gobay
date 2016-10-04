<?php

$tmpl = "
    o.AddFIELD()

    if o.FIELD.Include != true {
        t.Errorf(\"Failed, including FIELD is required %+v\",o.FIELD)
        return
    }

    txt, err = compileGoString(\"Test\", GetMyeBaySellingTemplate(), o, nil)
    if err != nil {
        t.Errorf(\"Could not compile file %s\",err)
        return
    }

    txt = fmt.Sprintf(\"<?xml version=\\\"1.0\\\" encoding=\\\"utf-8\\\"?><root>%s</root>\",txt)

    to = GetMyeBaySellingStructFromXML([]byte(txt))

    if to.FIELD.Include != o.FIELD.Include {
        t.Errorf(\"Failed to reload output\")
        return
    }
";

$ins = array("ActiveList",           
"SoldList",             
"UnsoldList",           
"BidList",              
"DeletedFromSoldList",  
"DeletedFromUnsoldList",
"ScheduledList"); 

foreach ($ins as $in ) {
    echo str_replace("FIELD",$in, $tmpl);
}       