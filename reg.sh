docker run --name registry \
  -p 8282:5000 -v /opt/registry/data:/var/lib/registry:z      \
  -d docker.io/library/registry:2
