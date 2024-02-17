package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Anjay Lani", "Lani"))
	fmt.Println(strings.Split("Anjay Lani", " "))
	fmt.Println(strings.ToLower("Anjay Lani"))
	fmt.Println(strings.ToUpper("Anjay Lani"))
	fmt.Println(strings.Trim("                Anjay Lani                ", " "))
	fmt.Println(strings.ReplaceAll("Adamma dadaodao Anjay Lani Anjay", "Anjay", "Ganti"))
}
