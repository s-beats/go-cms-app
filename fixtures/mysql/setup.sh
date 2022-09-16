#!bin/bash

function createTables() {
    DIR=$1
    FILE=$2

    mysql -u root go-cms < ${DIR}/ddl/${FILE}.ddl
}

function createDummyData() {
    DIR=$1

    DML=`find ${DIR}/dml -name "*.dml"`
    for D in ${DML}
    do
        echo Insert dml file: `basename ${D}`
        mysql -u root go-cms < ${D}
    done
}

MYSQL_PWD=password
DIR=`dirname $0`

export MYSQL_PWD

# TODO: ユーザー作成
# TODO: Database作成

createTables ${DIR} master
createTables ${DIR} transaction

createDummyData ${DIR}
