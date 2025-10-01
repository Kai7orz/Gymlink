#!/bin/sh

CMD_MYSQL="mysql -u${MYSQL_USER} -p${MYSQL_PASSWORD} ${MYSQL_DATABASE}"
$CMD_MYSQL -e "create table test ( 
    id int(10)  AUTO_INCREMENT NOT NULL primary key,
    title varchar(50) NOT NULL
    );"
$CMD_MYSQL -e  "insert into test values (1, 'テスト1');"
$CMD_MYSQL -e  "insert into test values (2, 'テスト2');"