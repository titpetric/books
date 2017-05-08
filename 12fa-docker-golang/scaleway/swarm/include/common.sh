function scw {
	if [ ! -f ".scwrc" ]; then
		echo '{}' > .scwrc && chmod 600 .scwrc
		docker run -it --rm --volume=$(pwd)/.scwrc:/.scwrc scaleway/cli login
	fi
	docker run -it --rm --volume=$(pwd)/.scwrc:/.scwrc scaleway/cli "$@"
}