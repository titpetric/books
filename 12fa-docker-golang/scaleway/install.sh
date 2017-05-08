#!/bin/bash
ADDR=$(./get-ip.sh)
scp scripts/* $ADDR:/root
ssh $ADDR /root/bootstrap.sh