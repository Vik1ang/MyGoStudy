package main

import "fmt"

type Animal interface {
	Eat()
	Run()
}

type Cat struct {
	Name string
	Sex  bool
}

type Dog struct {
	Name string
}

func (c Cat) Run() {
	fmt.Println(c.Name, "开始跑")
}

func (c Cat) Eat() {
	fmt.Println(c.Name, "开始吃")
}

func (d Dog) Run() {
	fmt.Println(d.Name, "开始跑")
}

func (d Dog) Eat() {
	fmt.Println(d.Name, "开始吃")
}

func main() {

	var a1 Animal
	c1 := Cat{
		Name: "Tom",
		Sex:  false,
	}
	a1 = c1
	a1.Run()
	a1.Eat()

	fmt.Println("================================")

	var a2 Animal
	a2 = Cat{
		Name: "Tom",
		Sex:  false,
	}
	a2.Run()
	a2.Eat()

	fmt.Println("================================")

	myFun1([]string{"123", "123"})

	fmt.Println("================================")

	c3 := Cat{
		Name: "Tom",
		Sex:  false,
	}
	myFun2(c3)

}

func myFun1(a interface{}) {
	fmt.Println(a)
}

func myFun2(a Animal) {
	a.Run()
	a.Eat()
}
