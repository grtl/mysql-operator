#!/bin/bash
docker login --email=$DOCKER_HUB_EMAIL --username=$DOCKER_HUB_USERNAME --password=$DOCKER_HUB_PASSWORD
docker push grtl/mysql-operator:latest
