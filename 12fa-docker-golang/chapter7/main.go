package main

import (
	"fmt"
	"log"
	"net/http"

	"app/api"
	"app/common"

	"github.com/namsral/flag"
)

func main() {
	// set up flags
	var (
		port            = flag.Int("port", 3000, "Listen port for server")
		nodeAppInstance = flag.Int("node-app-instance", 0, "PM2 application instance")
		redis           = flag.String("redis", "redis:6379", "Redis address (host:port)")
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
	log.Printf("Ready and listening on port %d + %d\n", *port, *nodeAppInstance)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", *port+*nodeAppInstance), nil); err != nil {
		panic(err)
	}
}
