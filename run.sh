#!/bin/bash

#cd $GOPATH/src/github.com/myapp
cd $GOPATH/src/github.com/authrecipe

buffalo plugins install github.com/gobuffalo/buffalo-pop

buffalo db create -a

buffalo db migrate

npm install

ADDR=0.0.0.0 buffalo dev

while true
do
    sleep 60
done
