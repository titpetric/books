#!/bin/bash
cd $(dirname $(readlink -f $0))
if [ ! -d "build/public_html" ]; then
	mkdir -p build/public_html
fi
rsync -a public_html/ build/public_html/
docker run --net=party --rm -it -p 8080:8080 -v `pwd`:/go/src/app -w /go/src/app golang go build -o build/gotwitter main.go
rsync -a build/ ../chapter8/app/
