package main

import "fmt"

func main() {
	a := 0
	a--
	// 没有前自增
	fmt.Println(a)
	fmt.Println("===================================")

	a = 1

	if a > 3 {
		fmt.Println("> 3")
	} else if a == 1 {
		fmt.Println(" = 1")
	} else {
		fmt.Println(" <= 3")
	}

	fmt.Println("===================================")

	switch a {
	case 0:
		fmt.Println("0")
	case 1:
		fmt.Println("1")
		fallthrough
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("here")
	}

	fmt.Println("===================================")

	a = 0
	for {
		a++
		fmt.Println(a)
		if a > 10 {
			break
		}
	}

	fmt.Println("===================================")

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("===================================")
	fmt.Println("goto")

	a = 0
A:
	for a < 10 {
		a++
		fmt.Println(a)
		if a == 5 {
			break A
			goto B
		}
	}

B:
	fmt.Println("rush B")
}
