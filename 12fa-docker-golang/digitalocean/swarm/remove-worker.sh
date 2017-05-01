#!/bin/bash
NODE=($(./list-workers.sh | head -n1))
NODE_NAME=${NODE[0]}
NODE_ADDR=${NODE[1]}
if [ ! -z "${NODE_ADDR}" ]; then
	echo "Leaving swarm: $NODE"
	ssh $NODE_ADDR docker node update --availability drain $NODE_NAME
	sleep 2
	ssh $NODE_ADDR docker node rm $NODE_NAME -f
	echo "Purging droplet: $NODE_NAME"
	doctl compute droplet delete $NODE_NAME -f -v
else
	echo "No workers to remove"
fi