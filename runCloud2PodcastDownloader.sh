#!/bin/bash
set -o nounset
set -o errexit

if [[ "$#" -ne 1 ]]; then
    echo "Usage: $(basename "$0") downloadDirectory"
    echo "Docker Container with name cloud2podcast will be started with the downloadDirectory for your podcasts mounted"
    exit 1
fi

echo "_____->Running cloud2podcastdownloader Application<-_____"
downloadDir=$1
docker run --rm --name c2pDL -v $downloadDir:/downloads -d cloud2podcastdownloader