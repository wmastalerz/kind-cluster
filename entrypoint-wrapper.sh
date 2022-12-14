#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

reg_name=registry
reg_port=8282
reg_ip=$(ip a |grep -A +3 eth0 |grep inet |awk '{split($2,a,"/"); print a[1]}' |awk '{split($1,a,"."); print a[1]"."a[2]"."a[3]".1"}' )

echo "Setting up KIND cluster"

# Startup docker.
echo "${reg_ip} ${reg_name}" >> /etc/hosts
mkdir -p /etc/docker && echo "{ \"insecure-registries\" : [\"${reg_name}:${reg_port}\"] }" > /etc/docker/deamon.json
service docker start

# Startup sshd.
ssh-keygen -t rsa -N "" -f ~/.ssh/id_rsa
cat ~/.ssh/id_rsa.pub >> ~/.ssh/authorized_keys
chmod og-wx ~/.ssh/authorized_keys
service ssh start
ssh -o StrictHostKeyChecking=no localhost exit

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
docker exec -it kind-control-plane sh -c "echo ${reg_ip} registry >> /etc/hosts"
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

# create currnet ns
kubectl create ns test
kubectl config set-context --current --namespace test

# add ip of registry into dns
kubectl get cm coredns -n kube-system -o jsonpath='{.data.*}' > Corefile
echo "${reg_ip} registry" > customdomains.db
kubectl create configmap coredns --from-file Corefile,customdomains.db -o yaml --dry-run | kubectl apply -n kube-system -f -

# set default python3
update-alternatives --install /usr/bin/python python /usr/bin/python3.7 1

exec "$@"
