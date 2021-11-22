package main

import (
	"GoBasicStudy/src/study01/dog"
	"GoBasicStudy/src/study01/testPackage"
	"fmt"
)

func main() {
	var a string = "hello 1010"
	// 关键字 变量名 变量类型 = 变量值
	fmt.Println(a)

	b := "hello 2020"
	// 隐式
	fmt.Println(b)

	// 关键字不能做变量名

	// test package
	fmt.Println(testPackage.A)
	fmt.Println(testPackage.B)

	// exercise
	fmt.Println(dog.Name)
}
