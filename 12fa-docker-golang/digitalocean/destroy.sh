#!/bin/bash
CHECK=$(doctl compute droplet list asx --format ID --no-header | wc -l)
if [ "$CHECK" == "0" ]; then
	echo "Droplet ASX not started"
else
	echo "Deleting droplet asx"
	doctl compute droplet delete asx -f -v
fi