#!/usr/bin/env sh

set -e
set -v
docker tag mysql-operator:testing grtl/mysql-operator:latest
docker login --username=$DOCKER_HUB_USERNAME --password=$DOCKER_HUB_PASSWORD
docker push grtl/mysql-operator:latest
