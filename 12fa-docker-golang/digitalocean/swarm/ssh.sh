#!/bin/bash
IFS=$'\n' NODES=$(./list-managers.sh)
for NODE in $NODES; do
	IFS=' ' NODE=($NODE)
	echo "> ${NODE[0]}"
	ssh ${NODE[1]} "$@"
done