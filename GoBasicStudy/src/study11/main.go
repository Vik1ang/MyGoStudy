package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age  int
	Sex  bool
}

type Student struct {
	Class string
	User
}

func (u User) SayName(name string) {
	fmt.Println("My name is", name)
}

func main() {
	u1 := User{
		Name: "name1",
		Age:  18,
		Sex:  true,
	}

	s1 := Student{
		Class: "Class",
		User:  User{},
	}

	check1(u1)

	fmt.Println("================================")

	check2(u1)
	check2(s1)

	fmt.Println("================================")

	check3(u1)

	fmt.Println("================================")

	u2 := User{
		Name: "name1",
		Age:  18,
		Sex:  true,
	}

	s2 := Student{
		Class: "Class",
		User:  u2,
	}

	check4(s2)

	fmt.Println("================================")

	check5(&s2)
	fmt.Println(s2)

	fmt.Println("================================")

}

func check1(v interface{}) {
	v.(User).SayName(v.(User).Name)
}

func check2(v interface{}) {
	switch v.(type) {
	case User:
		fmt.Println(v.(User).Name)
		fmt.Println("I am User")
	case Student:
		fmt.Println(v.(Student).Class)
		fmt.Println("I am Student")
	}
}

func check3(value interface{}) {
	t := reflect.TypeOf(value)
	v := reflect.ValueOf(value)

	for i := 0; i < t.NumField(); i++ {
		fmt.Println(v.Field(i))
	}
}

func check4(value interface{}) {
	t := reflect.TypeOf(value)
	v := reflect.ValueOf(value)

	fmt.Println(t, v)
	fmt.Println(v.FieldByIndex([]int{1, 0}))
	fmt.Println(v.FieldByName("Class"))

	ty := t.Kind()
	if ty == reflect.Struct {
		fmt.Println("I am struct")
	}
}

func check5(value interface{}) {
	v := reflect.ValueOf(value)
	e := v.Elem()
	e.FieldByName("Class").SetString("Class123")
	fmt.Println(value)
}
