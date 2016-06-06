package main

import "foundations/bootstrap"
import "log"
import "fmt"

func main() {
	redis, err := bootstrap.GetRedis()
	if err != nil {
		log.Fatal("Fatal error: ", err)
	}
	fmt.Printf("[%.4f] Starting\n", bootstrap.Now())
	pong, err := redis.Do("PING")
	fmt.Printf("[%.4f] Response %s, err %#v\n", bootstrap.Now(), pong, err)
}
