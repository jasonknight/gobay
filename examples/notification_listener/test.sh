#!/bin/bash

curl -X POST -d "@test_data/test.xml" http://localhost:8000
cat /tmp/gobay_notifications.log
