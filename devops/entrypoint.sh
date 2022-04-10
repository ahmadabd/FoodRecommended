#!/bin/bash

EXECUTABLE=$1
APP=$2
CONFIG_FILE=$3
declare -a variables=(
    "MYSQL_DATABASE"
    "MYSQL_PASSWORD"
    "MYSQL_USER"
    "MYSQL_HOST"
    "MYSQL_PORT"
)
for var in "${variables[@]}"
do
    sed -i "s|\$$var|${!var}|g" $CONFIG_FILE
done

cp $CONFIG_FILE config.yml

$EXECUTABLE $APP