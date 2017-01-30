package main

import (
	"app/api"
	"app/common"
	"flag"
	"log"
	"net/http"
)

func main() {
	// set up flags
	var (
		port  = flag.String("port", "8080", "Listen port for server")
		redis = flag.String("redis", "redis:6379", "Redis address (host:port)")
	)
	flag.Parse()

	// set up config
	r := common.NewRedis(common.RedisAddress(*redis))
	r.Save()

	// set up twitter api
	apiTwitter := api.Twitter{}
	apiTwitter.Register()

	// serve static assets
	http.Handle("/", http.FileServer(http.Dir("./public_html")))

	// start up http server
	log.Printf("Ready and listening on port " + *port + "\n")
	if err := http.ListenAndServe(":"+*port, nil); err != nil {
		panic(err)
	}
}
