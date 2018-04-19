#!/bin/bash
source ../shell/common.sh
dep ensure -v
go run *.go
