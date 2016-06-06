package main

import "foundations/bootstrap"
import "fmt"

func main() {
	redis := bootstrap.RedigoPool()
	defer redis.Close()

	for i := 0; i < 5; i++ {
		bootstrap.RedigoDo("PING")
	}

	fmt.Printf("[%.4f] Start\n", bootstrap.Now())

	sleep1_chan := make(chan string, 1)
	sleep2_chan := make(chan string, 1)

	go func() {
		fmt.Printf("[%.4f] Run sleep 100ms\n", bootstrap.Now())
		sleep1, err := bootstrap.RedigoDo("DEBUG", "SLEEP", "0.1")
		if err != nil {
			sleep1 = "ERROR"
		}
		sleep1_chan <- sleep1.(string)
	}()

	go func() {
		fmt.Printf("[%.4f] Run sleep 200ms\n", bootstrap.Now())
		sleep2, err := bootstrap.RedigoDo("DEBUG", "SLEEP", "0.2")
		if err != nil {
			sleep2 = "ERROR"
		}
		sleep2_chan <- sleep2.(string)
	}()

	var result string
	result = <-sleep1_chan
	fmt.Printf("[%.4f] End Sleep 100ms, result %s\n", bootstrap.Now(), result)
	result = <-sleep2_chan
	fmt.Printf("[%.4f] End Sleep 200ms, result %s\n", bootstrap.Now(), result)
}
