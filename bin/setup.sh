#!/bin/sh
#load env
if [ -f .env ]
then
  export $(cat .env | sed 's/#.*//g' | xargs)
fi


# run gomod
go mod init

# run go mod tidy
go mod tidy

# run repository test
cd repository
go test -v

# run usecase test
cd ..
cd usecase
go test -v

# run command test
cd ..
cd cmd
go test -v

# run create database and run migration
export MYSQL_PWD=$DB_PASS;
mysql -u $DB_USER \
-e "CREATE DATABASE IF NOT EXISTS $DB_NAME;";

# migrate table question
mysql -u $DB_USER $DB_NAME < ../database/migration.sql

# build app
cd ..
go build -o bin/quiz_master