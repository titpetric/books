package apiservice

import "fmt"
import "github.com/namsral/flag"

var (
	networkPort = flag.String("port", "8080", "Network port to listen on")
)

func HelloWorld() {
	flag.Parse();
	fmt.Printf("Hello world! Network port: %s\n", *networkPort);
}
