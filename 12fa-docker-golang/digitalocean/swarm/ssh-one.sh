#!/bin/bash
NODE=($(./list-managers.sh | head -n1))
echo "> ${NODE[0]}"
ssh ${NODE[1]} "$@"
