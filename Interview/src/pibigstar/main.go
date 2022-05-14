package main

import "fmt"

func main() {
	//q1()
	//q2()
	//q3()
	q4()
}

func q1() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	s1 := slice[2:5]
	s2 := s1[2:6:7]

	s2 = append(s2, 100)
	s2 = append(s2, 200)

	s1[2] = 20

	fmt.Println(slice)
	fmt.Println(s1)
	fmt.Println(s2)
}

func q2() {
	s := []int{1, 1, 1}
	fQ2(s)
	fmt.Println(s)
}

func fQ2(s []int) {
	for i := range s {
		s[i] += 1
	}
}

func q3() {
	s := []int{1, 1, 1}
	newS := q3MyAppend(s)

	fmt.Println(s)
	fmt.Println(newS)

	fmt.Println("########")

	s = newS
	fmt.Println(s)
	q3MyAppendPtr(&s)

	fmt.Println(s)
}

func q3MyAppend(s []int) []int {
	s = append(s, 100)
	return s
}

func q3MyAppendPtr(s *[]int) {
	*s = append(*s, 100)
	return
}

func q4() {
	ageMap := make(map[string]int)
	ageMap["qcrao"] = 18

	age1 := ageMap["stefno"]
	fmt.Println(age1)

	age2, ok := ageMap["stefno"]
	fmt.Println(age2, ok)
}
