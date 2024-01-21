package main

import "fmt"

func main() {
	firstname := "Anjay"
	lastname := "Baru"

	fmt.Println("Hello '", firstname, lastname, "'")

	fmt.Printf("Hello '%s %s '\n", firstname, lastname)
}
