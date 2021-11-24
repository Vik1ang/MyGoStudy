package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Printf("%T\n", 123.3)

	var string1 string
	string1 = "123"
	fmt.Printf("%T\n", string1)
	int111, _ := strconv.Atoi(string1)
	fmt.Println(int111)
	fmt.Printf("%T\n", int111)

}
