#!/bin/bash
HOSTS=$(./list-workers.sh | awk '{print $1}' ; ./list-managers.sh | awk '{print $1}')
for HOST in $HOSTS; do
	echo "Deleting $HOST"
	doctl compute droplet delete $HOST -f -v
done
if [ -z "$HOSTS" ]; then
	echo "Nothing to do"
fi