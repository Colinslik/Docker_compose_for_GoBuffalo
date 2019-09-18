#!/bin/bash

cd $GOPATH/src/github.com/coke

ADDR=0.0.0.0 buffalo dev

while true
do
    sleep 60
done
