function go {
	printenv | grep "GO_" > /tmp/docker.env
	echo "== go" "$@" "=="
	if [ -f "docker.args" ]; then
	        ARGS=$(cat docker.args | xargs echo -n)
	fi
	SPECIALS="tool get build fmt"
	for SPECIAL in $SPECIALS; do
	        if [ "$1" == "$SPECIAL" ]; then
	                ARGS=""
	        fi
	done
	if [ -z "$APP" ]; then
	        APP="app"
	fi
	WORKDIR="/go/src/$APP"
	docker run $ARGS --rm=true --env-file /tmp/docker.env -i -v $(pwd):$WORKDIR -w $WORKDIR golang go "$@"
	echo ""
}

function gvt {
	if [ $1 == "fetch" ]; then
		BASE="vendor/"
		if [ -d "$BASE$2" ]; then
			return
		fi
	fi
	docker run --dns=8.8.8.8 --dns=8.8.4.4 --rm=true -i -v $(pwd):/go/src justincormack/gvt "$@"
}
