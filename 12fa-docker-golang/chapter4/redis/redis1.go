package main

import (
	"flag"
	"fmt"
	"os"

	"app/service"
)

func main() {
	// set up flags
	flags := flag.NewFlagSet("default", flag.ContinueOnError)

	// set up redis service
	redis := &service.Redis{}
	redis.Flags("redis", "redis1:6379", "Redis DNS", flags)

	// parse flags
	flags.Parse(os.Args[1:])

	fmt.Printf("[%.4fms] Starting\n", service.Now())
	fmt.Printf("[%.4fms] Sending PING\n", service.Now())
	pong, err := redis.Do("PING")
	fmt.Printf("[%.4fms] Response %s, err %#v\n", service.Now(), pong, err)
}
