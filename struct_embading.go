package struct_embade

import "fmt"

type person struct {
	name string
	age  int
}

func (p person) about() {
	fmt.Printf("-----------------------------------\n")
	fmt.Printf("name: %s \n age:%d\n", p.name, p.age)
}

type employee struct {
	person //here we are directly using person struct indide employee struct it is anynomous too it is like same
	//name same type we can give name here but is is not anynomous then in e1 we need to access like
	//e1.p1.name(not too much readable)
	eid    string
	salary float64
}

func main() {

	e1 := employee{
		person: person{
			name: "amrit poudel",
			age:  27,
		},
		eid:    "AM123POU#1",
		salary: 10000000,
	}

	fmt.Println("name: ", e1.name)
	fmt.Println("age: ", e1.age)
	fmt.Println("employee id : ", e1.eid)
	fmt.Println("salary: ", e1.salary)
	fmt.Println("all: ", e1)
	e1.about()

}
