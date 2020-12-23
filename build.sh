#!/bin/bash

# variables
NAME="go-myip"
IMAGE=$NAME
VERSION="$(cat VERSION)"

# stop and remove
docker stop $NAME; docker rm $NAME

# build
docker build -t $IMAGE:$VERSION .

# run 
docker run -d -p 3020:5000 \
  --name $NAME $IMAGE:$VERSION
