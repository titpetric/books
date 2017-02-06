#!/bin/bash
cd $(dirname $(readlink -f $0))
rm -rf build
mkdir -p build/public_html
rsync -a public_html/ build/public_html/
echo "Building app"
docker run --net=party --rm -it -p 8080:8080 -v `pwd`:/go/src/app -w /go/src/app golang:1.8-alpine go build -o build/gotwitter main.go
rsync -a build/ ../chapter8/app/
