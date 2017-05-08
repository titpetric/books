#!/bin/bash
export DEBIAN_FRONTEND=noninteractive

# update APT caches
apt-get update

# install required pkgs
apt-get -qyy install \
    apt-transport-https \
    ca-certificates \
    curl \
    software-properties-common

# install docker repos

ARCH="armhf"
if [ "$(uname -m)" == "x86_64" ]; then
	ARCH="amd64"
fi

add-apt-repository \
   "deb [arch=$ARCH] https://download.docker.com/linux/ubuntu \
   $(lsb_release -cs) \
   stable"
add-apt-repository \
   "deb [arch=$ARCH] https://download.docker.com/linux/ubuntu \
   $(lsb_release -cs) \
   edge"

# install public key
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | apt-key add -

# install docker
apt-get update && apt-get -qyy install docker-ce