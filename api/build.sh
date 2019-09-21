#!/bin/bash

WorkDIR=$PWD


function start() {

    cd $WorkDIR/mydb

    docker-compose up -d

    cd $WorkDIR/myapi

    docker-compose up -d
}

function stop() {

    cd $WorkDIR/myapi

    docker-compose down

    cd $WorkDIR/mydb

    docker-compose down
}


case ${1} in

"start")
        start
	;;
"stop")
	stop
	exit 1
	;;
*)
        echo "Usage: $0 {start | stop}"
        exit 1
esac

