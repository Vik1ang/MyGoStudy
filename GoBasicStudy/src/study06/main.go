package main

import "fmt"

func main() {
	fmt.Println("================================")

	fun1(3, "> 1")

	fmt.Println("================================")

	r1, r2 := fun2(0, "data2")
	fmt.Println(r1, r2)

	fmt.Println("================================")

	r3, r4 := fun3(0, "data2")
	fmt.Println(r3, r4)

	fmt.Println("================================")
	fun4 := func(data1 string) {
		fmt.Println(data1)
	}
	fun4("匿名函数")

	fmt.Println("================================")

	fun5(9527, "1", "2", "3", "4")

	fmt.Println("================================")

	a1 := []string{"1", "2", "3", "4"}
	fun5(9527, a1...)

	fmt.Println("================================")

	fun6()(4)

	fmt.Println("================================")

	defer fun7()
	fmt.Println("1")
	fmt.Println("2")
	fmt.Println("3")

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

func fun3(data1 int, data2 string) (ret1 int, ret2 string) {
	ret1 = data1
	ret2 = data2
	return
}

func fun5(data1 int, data2 ...string) {
	fmt.Println(data1, data2)
	for k, v := range data2 {
		fmt.Println(k, v)
	}
}

func fun6() func(int) {
	return func(num int) {
		fmt.Println("闭包函数", num)
	}
}

func fun7() {
	fmt.Println("我想最先执行")
}
