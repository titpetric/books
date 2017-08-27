#!/bin/bash
CHECK=$(doctl compute droplet list asx --format ID --no-header | wc -l)
if [ "$CHECK" == "0" ]; then
	echo "Creating ASX droplet"
	doctl compute droplet create asx -v \
		--image docker-16-04 \
		--size 2gb \
		--region ams3 \
		--ssh-keys $(./ssh-key.sh)
else
	echo "Droplet ASX already running"
fi