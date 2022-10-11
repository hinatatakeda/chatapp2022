#!/bin/sh

CMD_MYSQL="mysql -u${MYSQL_USER} -p${MYSQL_PASSWORD} ${MYSQL_DATABASE}"

$CMD_MYSQL -e "create table user(
    id INT(10) NOT NULL AUTO_INCREMENT,
    username    varchar(50) NOT NULL,
    password    varchar(50) NOT NULL,
    email       varchar(50) NOT NULL,
    PRIMARY KEY (`id`)
    );"

$CMD_MYSQL -e  "insert into user values ('hinata', 'hinatahinata','hinata@gmail.com');"
$CMD_MYSQL -e  "insert into user values ('conan', 'hinatahinata2','takeda@gmail.com');"

$CMD_MYSQL -e "create table room(
    roomid      int(10)  AUTO_INCREMENT NOT NULL primary key,
    roomname    varchar(50) NOT NULL
    );"

$CMD_MYSQL -e  "insert into room values ('general');"
$CMD_MYSQL -e  "insert into room values ('random');"


$CMD_MYSQL -e "create table chat(
    chatid     int(10)  AUTO_INCREMENT NOT NULL primary key,
    roomid     varchar(50) NOT NULL,
    id int(10)   NOT NULL,
    text varchar(200) NOT NULL
    retext boolean NOT NULL
    );"


