#!/bin/bash

echo "_____->Building cloud2podcastdownloader Application<-_____"
docker build -t cloud2podcastdownloader -f DockerfileRPI .
