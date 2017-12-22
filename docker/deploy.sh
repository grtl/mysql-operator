#!/bin/bash
docker login --username=$DOCKER_HUB_USERNAME --password=$DOCKER_HUB_PASSWORD
docker push grtl/mysql-operator:latest
