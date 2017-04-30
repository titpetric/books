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
	redis.Flags("redis1", "redis1:6379", "Redis DNS", flags)
	redis.Flags("redis2", "redis2:6379", "Redis DNS", flags)
	redis.Flags("redis3", "redis3:6379", "Redis DNS", flags)

	// parse flags
	flags.Parse(os.Args[1:])

	fmt.Printf("[%.4f] Start\n", service.Now())

	sleep1_chan := make(chan string, 1)
	sleep2_chan := make(chan string, 1)

	go func() {
		conn := redis.Get()
		defer conn.Close()

		fmt.Printf("[%.4f] Run sleep 100ms\n", service.Now())
		sleep1, err := conn.Do("DEBUG", "SLEEP", "0.1")
		if err != nil {
			sleep1 = "ERROR"
		}
		sleep1_chan <- sleep1.(string)
	}()

	go func() {
		conn := redis.Get()
		defer conn.Close()

		fmt.Printf("[%.4f] Run sleep 200ms\n", service.Now())
		sleep2, err := conn.Do("DEBUG", "SLEEP", "0.2")
		if err != nil {
			sleep2 = "ERROR"
		}
		sleep2_chan <- sleep2.(string)
	}()

	var result string
	result = <-sleep1_chan
	fmt.Printf("[%.4f] End Sleep 100ms, result %s\n", service.Now(), result)
	result = <-sleep2_chan
	fmt.Printf("[%.4f] End Sleep 200ms, result %s\n", service.Now(), result)
}
