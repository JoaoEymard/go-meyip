#!/bin/bash

# variables
NAME="ws-juno"
IMAGE=$NAME
VERSION="$(cat VERSION)"

# stop and remove
docker stop $NAME; docker rm $NAME

# build
docker build -t $IMAGE:$VERSION .

# run 
docker run -d -p 3020:3000 \
  --name $NAME $IMAGE:$VERSION
