#!/bin/bash

cd $GOPATH/src/github.com/myapp

ADDR=0.0.0.0 buffalo dev

while true
do
    sleep 60
done
