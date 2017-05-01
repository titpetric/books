#!/bin/bash
doctl compute droplet list --tag-name swarm-worker --format Name,PublicIPv4 --no-header
