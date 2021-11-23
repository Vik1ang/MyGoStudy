package main

import "fmt"

func main() {
	fmt.Println("================================")
	
	fun1(3, "> 1")

	fmt.Println("================================")
	
	v1, v
	
}

func fun1(data1 int, data2 string) {
	if data1 > 1 {
		fmt.Println(data1)
	} else {
		fmt.Println(data2)
	}
}

func fun2(data1 int, data2 string) (ret1 int, ret2 string) {
	if data1 > 1 {
		return data1, ">1"
	} else {
		return data1, data2 + "<1"
	}
}