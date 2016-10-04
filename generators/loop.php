<?php

$tmpl = "
    o = NewMyeBaySellingStruct()
    o.Add\$in()
    if o.\$in.Include != true {
        t.Errorf(\"Failed, including \$in is required here %+v\",o.\$in)
        return
    }
    txt, err = compileGoString(\"Test\$in\", GetMyeBaySellingTemplate(), o, nil)
    if err != nil {
        t.Errorf(\"Could not compile file %s\", err)
        return
    }
    txt = fmt.Sprintf(\"<?xml version=\\\"1.0\\\" encoding=\\\"utf-8\\\"?><root>%s</root>\",txt)

    to = GetMyeBaySellingStructFromXML([]byte(txt))

    if to.\$in.Include != o.\$in.Include {
        t.Error(\"Failed to reload output for \$in [[%s]]\",txt)
        return
    }";

$ins = array("BidList","ScheduledList","SoldList","UnsoldList","DeletedFromUnsoldList");
foreach ($ins as $in) {
    $str = str_replace('$in',$in,$tmpl);
    echo $str . "\n";
}