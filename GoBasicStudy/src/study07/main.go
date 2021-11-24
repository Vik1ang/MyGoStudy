package main

import "fmt"

func main() {
	var a1 int
	a1 = 123
	fmt.Println(a1)
	var b1 int
	b1 = a1
	fmt.Println(a1, b1)
	b1 = 456
	fmt.Println(a1, b1)

	fmt.Println("================================")

	var a2 int
	a2 = 123
	fmt.Println(a2)
	var b2 *int
	b2 = &a2
	fmt.Println(a2, b2)
	*b2 = 999
	fmt.Println(a2, b2, *b2)

	fmt.Println("================================")

	// 数组指针
	var arr1 [5]string
	arr1 = [5]string{"1", "2", "3", "4", "5"}
	fmt.Println(arr1)
	var arr1P *[5]string
	arr1P = &arr1
	fmt.Println(arr1, arr1P)

	// 指针数组
	var arrP1 [5]*string
	var str1 = "str1"
	var str2 = "str2"
	var str3 = "str3"
	var str4 = "str4"
	var str5 = "str5"

	arrP1 = [5]*string{&str1, &str2, &str3, &str4, &str5}
	fmt.Println(arrP1, str3)
	*arrP1[2] = "555"
	fmt.Println(arrP1, str3)

	fmt.Println("================================")
	var v1 = "我定义了"
	fmt.Println(v1)
	fun1(&v1)
	fmt.Println(v1)

	fmt.Println("================================")

	var v2 = "我是来测地址的"
	p2 := &v2
	*p2 = "我是来测地址的123123"
	fmt.Println(v2)
}

func fun1(p1 *string) {
	*p1 = "我变了"
}
