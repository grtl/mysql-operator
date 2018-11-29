#!/usr/bin/env sh

set -e
set -v
curl -Lo kubectl https://storage.googleapis.com/kubernetes-release/release/$K8S_VERSION/bin/linux/amd64/kubectl \
    && chmod +x kubectl && sudo mv kubectl /usr/local/bin/
curl -Lo minikube https://storage.googleapis.com/minikube/releases/$MINIKUBE_VERSION/minikube-linux-amd64 \
    && chmod +x minikube && sudo mv minikube /usr/local/bin/
