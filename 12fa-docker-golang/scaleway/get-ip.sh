#!/bin/bash
source swarm/include/common.sh
scw inspect asx | jq -r '.[0].public_ip.address'