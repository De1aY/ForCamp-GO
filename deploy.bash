#!/bin/bash
while getopts u:r:p: option
do
case "${option}"
in
u) USER=${OPTARG};;
p) PASSWORD=${OPTARG};;
r) REGISTRY=${OPTARG};;
esac
done
docker login -u $USER -p $PASSWORD $REGISTRY
if docker ps | grep -q wplay_backend
then
docker kill -s HUP wplay_backend
fi
cd /usr/app/source
docker-compose pull
docker-compose up -d