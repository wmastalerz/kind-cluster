#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

reg_name=registry
reg_port=8282


echo "Setting up KIND cluster"

# Startup docker.
echo "172.17.0.1 ${reg_name}" >> /etc/hosts
mkdir -p /etc/docker && echo "{ \"insecure-registries\" : [\"${reg_name}:${reg_port}\"] }" > /etc/docker/deamon.json
service docker start

# Startup kind.
cat <<EOF | kind create cluster --image=${KIND_NODE_IMAGE-"kindest/node:v1.21.1"} --config=-
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
containerdConfigPatches:
- |-
  [plugins."io.containerd.grpc.v1.cri".registry.mirrors."${reg_name}:${reg_port}"]
    endpoint = ["http://${reg_name}:${reg_port}"]
EOF


#kind create cluster --config=kind-config.yaml --image=${KIND_NODE_IMAGE-"kindest/node:v1.21.0"} --wait=900s
docker exec -it kind-control-plane sh -c 'echo "172.17.0.1 registry" >> /etc/hosts'
docker cp /usr/share/zoneinfo kind-control-plane:/usr/share/zoneinfo/
for node in $(kind get nodes); do
  kubectl annotate node "${node}" "kind.x-k8s.io/registry=localhost:${reg_port}";
done
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: ConfigMap
metadata:
  name: local-registry-hosting
  namespace: kube-public
data:
  localRegistryHosting.v1: |
    host: "${reg_name}:${reg_port}"
    help: "https://kind.sigs.k8s.io/docs/user/local-registry/"
EOF
kubectl create ns csdb
kubectl config set-context --current --namespace csdb
exec "$@"
