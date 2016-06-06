package main

import "foundations/bootstrap"
import "log"
import "fmt"

func main() {
	redis, err := bootstrap.GetRedis()
	if err != nil {
		log.Fatal("Fatal error: ", err)
	}
	redis.Do("PING")

	fmt.Printf("[%.4f] Starting\n", bootstrap.Now())

	sleep1, err := redis.Do("DEBUG", "SLEEP", "0.1")
	fmt.Printf("[%.4f] End Sleep 100ms, result %s err %v\n",
		bootstrap.Now(), sleep1, err)

	sleep2, err := redis.Do("DEBUG", "SLEEP", "0.2")
	fmt.Printf("[%.4f] End Sleep 200ms, result %s err %v\n",
		bootstrap.Now(), sleep2, err)
}
