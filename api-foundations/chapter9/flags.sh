#!/bin/bash
source ../shell/common.sh
gvt fetch "github.com/namsral/flag"
printenv | grep GO_ > /tmp/docker.env
docker run --rm --env-file /tmp/docker.env -i -v `pwd`:/go/src/app -w /go/src/app golang go run flags.go
