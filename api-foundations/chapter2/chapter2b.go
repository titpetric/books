package main

import "app/bootstrap"
import "fmt"

func main() {
	fmt.Printf("Time: %.4f\n", bootstrap.Now())
	fmt.Printf("Hello world!\n")
	fmt.Printf("Time: %.4f\n", bootstrap.Now())

	bootstrap.StartTime = 0
	fmt.Printf("Time after reset: %.4f\n", bootstrap.Now())
}
