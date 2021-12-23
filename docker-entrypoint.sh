#!/bin/bash
if [ ! -f /app/db/PrometheusAlertDB.db ]; then
    cp /opt/PrometheusAlertDB.db /app/db/PrometheusAlertDB.db
    echo 'init ok!'
else
    echo 'pass!'
fi

if env | grep -q '^PA_[^=].\+'; then
    for VAR_NAME in $(env | grep '^PA_[^=].\+' | sed -r "s/^PA_([^=]*).*/\1/g"); do
        VAL_CONTENT=$(eval echo \${PA_${VAR_NAME}})
        echo "Config overridden from Environment variable, var=${VAR_NAME}=${VAL_CONTENT}"
        sed -i -E "s/(${VAR_NAME})(\ *=\ *).*/\1\2${VAL_CONTENT}/i" /app/conf/app.conf
    done
fi

exec /app/PrometheusAlert
