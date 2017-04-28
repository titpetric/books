package main

import (
	"os"
	"fmt"
	"flag"

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


	redis.Do("PING")
	fmt.Printf("[%.4f] Starting\n", service.Now())

	sleep1, err := redis.Do("DEBUG", "SLEEP", "0.1")
	fmt.Printf("[%.4f] End Sleep 100ms, result %s err %v\n",
		service.Now(), sleep1, err)

	sleep2, err := redis.Do("DEBUG", "SLEEP", "0.2")
	fmt.Printf("[%.4f] End Sleep 200ms, result %s err %v\n",
		service.Now(), sleep2, err)
}
