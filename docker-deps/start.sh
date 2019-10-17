#!/bin/bash
set -e

#env | while read i
#do
#    if [[ $i =~ ^VARNISH ]]; then
#        IFS='=' read -r param value <<< "$i"
#        sed  -i "s~${param}~${value}~" /etc/supervisor/supervisord.conf
#    fi
#done

exec /usr/bin/supervisord -c /etc/supervisor/supervisord.conf