# Find current network
ip=$(ip r g 8.8.8.8 |awk '{split($7,a,"/"); print a[1]}' |head -n +1)
dev=$(ip r g 8.8.8.8 |awk '{split($5,a,"/"); print a[1]}' |head -n +1)
subnet=${ip:0:10}0

# Step 0: Reset
{ # the redirect of this block also hides bash -x
    sudo ifdown --all
    for zUp in $(ip addr show | sed -nEe 's/[0-9]+: ([^:]+).* UP .*/\1/p' | tr '\n' ' '); do
        sudo ip link set "$zUp" down # force interface down which ifdown was unable to bring down
    done
    sudo iptables -F
    sudo iptables -t nat -F
    sudo systemctl stop dnsmasq.service
    sudo systemctl disable dnsmasq.service
} &> /dev/null
echo ip ${ip}/24 dev ${dev} subnet ${subnet}/24
sudo ip addr del ${ip}/24 dev ${dev}

# Step 1: Create
# ERROR: can't add wlp1s0 to bridge virbr-kind0: Operation not supported
# Create "shared_nw" with a bridge name "docker1"
sudo docker network create \
    --driver bridge \
    --subnet=${subnet}/24 \
    --gateway=${ip} \
    --opt "com.docker.network.bridge.name"="virbr-kind0" \
    kind

# Add kind to virbr-kind0 
sudo brctl addif virbr-kind0 ${dev}
sudo ifup --all
sudo ifup wlp1s0 2>&1 | grep -Ee "^Listening on" -e "^DHCPDISCOVER" -e "^bound to"
if [[ $(ip link show wlp1s0) = *' state DOWN '* ]]; then
    echo "$(tput setaf 1)!!! $(basename "$0"): Could not ifup wlNet !!!$(tput sgr0)"
  exit
fi
