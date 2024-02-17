package main

import (
	"fmt"
	"strconv"
)

func main() {

	// String to Boolean
	result, err := strconv.ParseBool("false")
	if err != nil {
		fmt.Println("Error ", err.Error())
	} else {
		fmt.Println(result)
	}

	resultInt, err := strconv.Atoi("12343")
	if err != nil {
		fmt.Println("Error ", err.Error())
	} else {
		fmt.Println(resultInt)
	}

	binary := strconv.FormatInt(999, 2)
	fmt.Println(binary)

	var string = strconv.Itoa(999)
	fmt.Println(string)
}
