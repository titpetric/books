#!/bin/bash
TAG="swarm"
SSH_KEY=$(./ssh-key.sh)

## create tag for docker swarm
doctl compute tag create $TAG

## see if tagged images are run yet
IPADDR=$(doctl compute droplet list --tag-name swarm --format PublicIPv4 --no-header | head -n1)
if [ ! -d "cloud-init" ]; then
	mkdir cloud-init
fi

## creating a swarm
SCRIPT_CREATE="#!/bin/bash
ufw allow 2377/tcp
export PUBLIC_IPV4=\$(curl -s http://169.254.169.254/metadata/v1/interfaces/public/0/ipv4/address)
docker swarm init --advertise-addr \"\${PUBLIC_IPV4}:2377\""

echo "$SCRIPT_CREATE" > cloud-init/create.sh

## joining a swarm
if [ ! -z "$IPADDR" ]; then
	TOKEN=$(ssh $IPADDR docker swarm join-token -q manager)
	if [ -z "$TOKEN" ]; then
		echo "Couldn't get swarm token from running node"
		exit 1
	fi
	SCRIPT_JOIN="#!/bin/bash
ufw allow 2377/tcp
export PUBLIC_IPV4=\$(curl -s http://169.254.169.254/metadata/v1/interfaces/public/0/ipv4/address)
docker swarm join --advertise-addr \"\${PUBLIC_IPV4}:2377\" --token \"$TOKEN\" \"$IPADDR:2377\""

	echo "$SCRIPT_JOIN" > cloud-init/join.sh
fi

chmod a+x cloud-init/*.sh

DROPLET="$TAG-$(date +%s)"
if [ -z "$IPADDR" ]; then
	## create a new swarm manager
	echo "Creating a new swarm host: $DROPLET (takes about a minute)"
	doctl compute droplet create $DROPLET -v --wait \
		--image docker-16-04 \
		--size 2gb \
		--tag-name $TAG \
		--enable-private-networking \
		--region ams3 \
		--ssh-keys ${SSH_KEY} \
		--user-data-file ./cloud-init/create.sh
else
	## join to existing swarm as manager
	echo "Adding a host to existing swarm: $DROPLET $IPADDR"
	doctl compute droplet create $DROPLET -v \
		--image docker-16-04 \
		--size 2gb \
		--tag-name $TAG \
		--enable-private-networking \
		--region ams3 \
		--ssh-keys ${SSH_KEY} \
		--user-data-file ./cloud-init/join.sh
fi