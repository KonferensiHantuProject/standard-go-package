package main

import (
	"flag"
	"fmt"
)

func main() {
	var string = flag.String("username", "root", "database username")
	var password = flag.String("password", "root", "database paawrod")
	var host = flag.String("host", "root", "database host")
	var port = flag.Int("port", 0, "database host")

	flag.Parse()

	fmt.Println("string", *string)
	fmt.Println("password", *password)
	fmt.Println("host", *host)
	fmt.Println("port", *port)

	// Example of how to run it
	//  go run flag.go -username=Addddduh -password="Adadada Asas" -host=localhost -port=1212121
}
