#!/usr/bin/env sh

set -e
curl -Lo kubectl https://github.com/kubernetes/kubernetes/releases/download/$K8S_VERSION/kubernetes.tar.gz \
    && chmod +x kubectl && sudo mv kubectl /usr/local/bin/
curl -Lo minikube https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64 \
    && chmod +x minikube && sudo mv minikube /usr/local/bin/