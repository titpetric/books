package main

import "github.com/namsral/flag"
import "fmt"
import "os"

func main() {
	fs := flag.NewFlagSetWithEnvPrefix(os.Args[0], "GO", 0)
	var (
		port = fs.Int("port", 8080, "Port number of service")
	)
	fs.Parse(os.Args[1:])

	fmt.Printf("Server port: %d\n", *port)
}
