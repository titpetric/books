package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/namsral/flag"
)

func main() {
	var age int

	godotenv.Load()

	flag.IntVar(&age, "age", 0, "age of gopher")
	flag.Parse()

	fmt.Print("age:", age)
}
