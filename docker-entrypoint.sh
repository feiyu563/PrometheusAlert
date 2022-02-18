#!/bin/bash
if [ ! -f /app/db/PrometheusAlertDB.db ]; then
    cp /opt/PrometheusAlertDB.db /app/db/PrometheusAlertDB.db
    echo 'init ok!'
else
    echo 'pass!'
fi

if env | grep -q '^PA_.\+=.\+'; then
    for VAR_NAME in $(env | grep '^PA_.\+=.\+' | sed -r "s/^PA_([^=]*).*/\1/g");do
        if echo ${VAR_NAME} | grep -q '.*\-.*'; then
            echo "\"PA_${VAR_NAME}\" in Environment variable contains '-',this will be ignored."
            continue
        fi
        CONF_ITEM=$(grep -Eio "${VAR_NAME/_/-}|${VAR_NAME}" /app/conf/app.conf)
        if [[ -z ${CONF_ITEM} ]]; then
            echo "\"PA_${VAR_NAME}\" in Environment variable not found from config file"
            continue
        fi
        CONF_CONTENT=$(eval echo \${PA_${VAR_NAME}})
        echo "Config overridden from Environment variable, ${CONF_ITEM}=${CONF_CONTENT}."
        sed -i -E "s@(${CONF_ITEM})(\ *=\ *).*@\1\2${CONF_CONTENT}@i" /app/conf/app.conf
    done
fi

exec /app/PrometheusAlert
