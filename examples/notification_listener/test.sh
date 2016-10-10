#!/bin/bash

curl -X POST -d "@test_data/test.xml" http://notifications.jolierouge.net:8000
cat /tmp/gobay_notifications.log
