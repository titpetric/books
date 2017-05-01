#!/bin/bash
CHECK=$(doctl compute droplet list asx --format ID --no-header | wc -l)
if [ "$CHECK" == "0" ]; then
	echo "Creating ASX droplet"
	doctl compute droplet create asx -v \
		--image docker-16-04 \
		--size 2gb \
		--region ams3 \
		--ssh-keys 01:23:45:67:89:0a:bc:de:fe:dc:ba:98:76:54:32:10
else
	echo "Droplet ASX already running"
fi