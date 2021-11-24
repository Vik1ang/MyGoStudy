package main

import "fmt"

type Qimiao struct {
	Name  string
	Age   int
	Sex   bool
	Hobby []string
	Home
}

type Home struct {
	P string
}

func (q *Qimiao) Song(name string) (ret string) {
	ret = "AAAA"
	fmt.Printf("%v唱了一首%v, 观众觉得%v\n", q.Name, name, ret)
	return
}

func (h *Home) Open() {
	fmt.Println("open", h.P)
}

func main() {
	var qm1 Qimiao
	qm1.Age = 18
	qm1.Name = "Name"
	qm1.Sex = true
	qm1.Hobby = []string{"a", "b"}
	fmt.Println(qm1)

	fmt.Println("================================")

	qm2 := Qimiao{
		Name:  "Name",
		Age:   18,
		Sex:   true,
		Hobby: []string{"a", "b"},
	}
	qmFunc1(qm2)

	fmt.Println("================================")

	qm3 := new(Qimiao)
	qm3.Name = "Name"
	qm3.Sex = true
	fmt.Println(qm3)

	fmt.Println("================================")

	var qm4P *Qimiao
	qm4P = &qm2
	qm4P.Name = "修改"
	// *qm4P.Name = "123"
	(*qm4P).Name = "修改123"
	fmt.Println(qm2, qm4P)

	fmt.Println("================================")

	qm5 := Qimiao{
		Name:  "Name",
		Age:   18,
		Sex:   true,
		Hobby: []string{"a", "b"},
	}
	qm5.Song("A")

	fmt.Println("================================")

	qm6 := Qimiao{
		Name:  "Name",
		Age:   18,
		Sex:   true,
		Hobby: []string{"a", "b"},
	}
	qm6.P = "S"
	fmt.Println(qm6)
	qm6.Open()

}

func qmFunc1(qm Qimiao) {
	fmt.Println(qm)
}
