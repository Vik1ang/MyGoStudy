package main

import "fmt"

func main() {
	var m1 map[string]string
	m1 = map[string]string{}
	m1["name"] = "m1"
	m1["sex"] = "male"
	fmt.Println(m1)

	fmt.Println("==========================")

	m2 := map[string]string{}
	m2["name"] = "name"
	m2["sex"] = "sex"
	fmt.Println(m2)

	fmt.Println("==========================")

	m3 := make(map[string]string)
	m3["name"] = "name"
	m3["sex"] = "sex"
	fmt.Println(m3)

	fmt.Println("==========================")

	m4 := map[int]bool{}
	m4[1] = true
	m4[2] = false
	fmt.Println(m4)

	fmt.Println("==========================")

	m5 := map[int]interface{}{}
	m5[1] = 1
	m5[2] = false
	m5[3] = "str"
	m5[4] = []int{1, 2, 3, 4}
	fmt.Println(m5, len(m5))

	// delete
	delete(m5, 4)
	fmt.Println(m5, len(m5))

	fmt.Println("==========================")

	m6 := map[string]interface{}{}
	m6["a"] = 1
	m6["b"] = false
	m6["c"] = "str"
	m6["d"] = []int{1, 2, 3, 4}

	for k, v := range m6 {
		fmt.Println(k, v)
	}
}
