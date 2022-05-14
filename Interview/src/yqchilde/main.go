package main

import "fmt"

func main() {
	//q1()
	//q2()
	//q3_1()
	//q3_2()
	//q9()
	//q11()
	//q12()
	//q14()
	//q17()
	//q18()
	//q19()
	//q23()
	q30()
}

func q1() {
	defer func() {
		fmt.Println("打印前")
	}()
	defer func() {
		fmt.Println("打印中")
	}()
	defer func() {
		fmt.Println("打印后")
	}()

	panic("触发异常")
}

func q2() {
	slice := []int{0, 1, 2, 3}

	for k, v := range slice {
		fmt.Println(k, v)
	}

	m := make(map[int]*int)

	for key, val := range slice {
		m[key] = &val
	}

	for k, v := range m {
		fmt.Println(k, "->", *v)
	}
}

func q3_1() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}

func q3_2() {
	s := make([]int, 0)
	s = append(s, 1, 2, 3, 4)
	fmt.Println(s)
}
func q9() {
	sn1 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}
	sn2 := struct {
		age  int
		name string
	}{age: 11, name: "11"}

	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}

	//sm1 := struct {
	//	age int
	//	m   map[string]string
	//}{age: 11, m: map[string]string{"a": "1"}}
	//sm2 := struct {
	//	age int
	//	m   map[string]string
	//}{age: 11, m: map[string]string{"a": "1"}}

	//if sm1 == sm2 {
	//	fmt.Println("sm1 == sm2")
	//}
}

type MyInt1 int
type MyInt2 = int

func q11() {
	var i int = 0
	//var i1 MyInt1 = i
	var i1 MyInt1 = MyInt1(i)
	var i2 MyInt2 = i
	fmt.Println(i1, i2)
}

func q12() {
	a := []int{7, 8, 9}
	fmt.Printf("%+v\n", a)
	q12Ap(a)
	fmt.Printf("%+v\n", a)
	q12App(a)
	fmt.Printf("%+v\n", a)
}
func q12Ap(a []int) {
	a = append(a, 10)
}
func q12App(a []int) {
	a[0] = 1
}

const (
	x = iota
	_
	y
	z = "zz"
	k
	p = iota
)

func q14() {
	fmt.Println(x, y, z, k, p)
}

func q17() {
	h := hello
	if h == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}
}

func hello() []string {
	return nil
}

type person struct {
	name string
}

func q18() {
	var m map[person]int
	p := person{"make"}
	fmt.Println(m[p])
}

func hello19(num ...int) {
	num[0] = 18
}
func q19() {
	i := []int{5, 6, 7}
	hello19(i...)
	fmt.Println(i[0])
}

func q22() {
	//a := 5
	//b := 8.1
	//fmt.Println(a + b)
}

func q23() {
	a := [5]int{1, 2, 3, 4, 5}
	t := a[3:4:4]
	fmt.Println(t[0])
}

func q24() {
	//a := [2]int{5, 6}
	//b := [3]int{5, 6}
	//if a == b {
	//
	//}
}

func q26() {
	var i interface{}
	if i == nil {
		fmt.Println("nil")
		return
	}
	fmt.Println("not nil")
}

func q27() {
	s := make(map[string]int)
	delete(s, "h")
	fmt.Println(s["h"])
}

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}

func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) showB() {
	fmt.Println("teacher showB")
}

func q30() {
	t := Teacher{}
	t.People.ShowB()
	t.ShowB()
}
