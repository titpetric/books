#!/bin/bash
doctl compute droplet list --tag-name swarm --format Name,PublicIPv4 --no-header
