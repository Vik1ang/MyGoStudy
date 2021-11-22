package main

import "fmt"

func main() {
	a1 := [3]int{0, 1, 2}
	fmt.Println(a1)

	fmt.Println("================")

	a2 := [...]int{1, 2, 3, 12, 323, 43, 43, 5, 342, 43, 324, 23, 423, 42, 32}
	fmt.Println(a2)

	fmt.Println("================")

	var a3 = new([10]int)
	a3[5] = 3
	fmt.Println(a3)

	fmt.Println("================")

	a4 := [...]string{"Dog", "Cat", "Monkey"}
	for i := 0; i < len(a4); i++ {
		fmt.Println(a4[i] + " Run")
	}

	for i, v := range a4 {
		fmt.Println(i, v)
	}

	fmt.Println("================")

	fmt.Println(len(a4), cap(a4))

	fmt.Println("================")

	a5 := [3][3]int{
		{0, 1, 2},
		{1, 2, 3},
		{2, 3, 4},
	}
	fmt.Println(a5)

	fmt.Println("================")

	// slice
	a6 := [3]int{0, 1, 2}
	c1 := a6[2:]
	fmt.Println(c1)
	c1[0] = 5
	fmt.Println(a6)

	fmt.Println("================")

	a7 := [3]int{0, 1, 2}
	c2 := a7[1:]
	fmt.Println(c2) // [ )
	c2 = append(c2, 5)
	c2[0] = 4
	fmt.Println(a7, c2)

	fmt.Println("================")

	a8 := [3]int{0, 1, 2}
	c3 := a8[:]
	fmt.Println(len(c3), cap(c3))
	c3 = append(c3, 5)
	fmt.Println(len(c3), cap(c3))
	c3 = append(c3, 5)
	fmt.Println(len(c3), cap(c3))
	c3 = append(c3, 5)
	fmt.Println(len(c3), cap(c3))
	c3 = append(c3, 5)
	fmt.Println(len(c3), cap(c3))
	c3 = append(c3, 5)
	fmt.Println(len(c3), cap(c3))

	fmt.Println("================")

	a9 := [3]int{0, 1, 2}
	c4 := a9[:]
	c5 := a9[2:]
	c4 = append(c4, 5)
	copy(c4[1:2], c5)
	fmt.Println(c4, c5)
}
