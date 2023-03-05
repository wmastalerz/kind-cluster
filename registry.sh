reg_name='registry'
reg_port='8282'
running="$(docker inspect -f '{{.State.Running}}' "${reg_name}" 2>/dev/null || true)"
if [ "${running}" != 'true' ]; then
  docker run -d -e BIND_ADDR=0.0.0.0:8282 --restart=always -p 127.0.0.1:8282:5000 --name "${reg_name}" registry:2
fi
