#!/usr/bin/env bash

cd /var/www/staging.imgtrip.com

if [ -e redeploy_flag.txt ];then
    bash -c "bash deploy.sh"
    rm -f redeploy_flag.txt
fi

cd /var/www/staging.x.imgtrip.com

if [ -e redeploy_flag.txt ];then
    bash -c "bash deploy.sh"
    rm -f redeploy_flag.txt
fi

cd /home/qskane/scripts && bash redeploy-staging.sh >> /dev/null 2>&1