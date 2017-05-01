#!/bin/bash
NODES=$(./list-managers.sh | wc -l)
if (( "$NODES" < 2 )); then
	echo "We need at least 2 managers to remove one manager"
	exit 1
fi

NODE=($(./list-managers.sh | head -n1))
NODE_NAME=${NODE[0]}
NODE_ADDR=${NODE[1]}
if [ ! -z "${NODE_ADDR}" ]; then
	echo "Leaving swarm: $NODE"
	ssh $NODE_ADDR docker node update --availability drain $NODE_NAME
	sleep 2
	ssh $NODE_ADDR docker node demote $NODE_NAME
	echo "Purging droplet: $NODE_NAME"
	doctl compute droplet delete $NODE_NAME -f -v
fi

MGR=$(./list-managers.sh | tail -n1 | awk '{print $2}')
ssh $MGR docker node rm $NODE_NAME -f
