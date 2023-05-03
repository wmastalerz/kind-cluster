#!/bin/sh
# Active network interface should be ethernet
# Linux known issue: not able to add wifi interface to bridge

set -o errexit

reg_name='registry'
reg_port='8282'
k8s_name='kind'

# Cleanup old kind cluster
kind delete cluster
docker stop ${reg_name} > /dev/null 2>&1 || true  && docker rm ${reg_name} > /dev/null 2>&1 || true
docker network rm ${k8s_name} > /dev/null 2>&1 || true

# Create bridge and kind network on external if
ip=$(ip r g 8.8.8.8 |awk '{split($7,a,"/"); print a[1]}' |head -n +1)
dev=$(ip r g 8.8.8.8 |awk '{split($5,a,"/"); print a[1]}' |head -n +1)
subnet=$(echo $ip |cut -c 1-10)0
echo ip ${ip}/24 dev ${dev} subnet ${subnet}/24
sudo ip addr del ${ip}/24 dev ${dev}
# Create "shared_nw" with a bridge name "docker1"
sudo docker network create \
    --driver bridge \
    --subnet=${subnet}/24 \
    --gateway=${ip} \
    --opt "com.docker.network.bridge.name"="virbr-kind0" \
    kind
# Add docker1 to eth1
sudo brctl addif virbr-kind0 ${dev}

# create registry container unless it already exists
if [ "$(docker inspect -f '{{.State.Running}}' "${reg_name}" 2>/dev/null || true)" != 'true' ]; then
  docker run \
    -d --restart=always -p "127.0.0.1:${reg_port}:5000" --name "${reg_name}" \
    registry:2
fi

#ip=${1:-"127.0.0.1"}
ip=$(ip r g 8.8.8.8 |awk '{split($7,a,"/"); print a[1]}' |head -n +1)
# create a cluster with the local registry enabled in containerd
cat <<EOF | kind create cluster --config=-
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
containerdConfigPatches:
- |-
  [plugins."io.containerd.grpc.v1.cri".registry.mirrors."localhost:${reg_port}"]
    endpoint = ["http://${reg_name}:5000"]
networking:
  # WARNING: It is _strongly_ recommended that you keep this the default
  # (127.0.0.1) for security reasons. However it is possible to change this.
  apiServerAddress: ${ip}
  # By default the API server listens on a random open port.
  # You may choose a specific port but probably don't need to in most cases.
  # Using a random port makes it easier to spin up multiple clusters.
  #apiServerPort: 6443
nodes:
- role: control-plane
- role: worker
EOF

# connect the registry to the cluster network if not already connected
if [ "$(docker inspect -f='{{json .NetworkSettings.Networks.kind}}' "${reg_name}")" = 'null' ]; then
  docker network connect "kind" "${reg_name}"
fi

# Document the local registry
# https://github.com/kubernetes/enhancements/tree/master/keps/sig-cluster-lifecycle/generic/1755-communicating-a-local-registry
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: ConfigMap
metadata:
  name: local-registry-hosting
  namespace: kube-public
data:
  localRegistryHosting.v1: |
    host: "localhost:${reg_port}"
    help: "https://kind.sigs.k8s.io/docs/user/local-registry/"
EOF


