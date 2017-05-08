#!/bin/bash
source swarm/include/common.sh
echo "Creating server asx"
scw _patch $(scw create --name=asx --commercial-type=X64-2GB ubuntu-xenial | sed 's/\-.*//g') tags="docker"
echo "Starting server asx"
scw start -w asx
ADDR=$(./get-ip.sh)
scp swarm/scripts/* $ADDR:/root
ssh $ADDR /root/bootstrap.sh
echo "Finished."