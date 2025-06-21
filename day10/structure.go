package struct1

import "fmt"

func main() {

	// type person struct {
	// 	fname string
	// 	lname string
	// 	age   int
	// }
	// // creating objects of struct
	// p1 := person{
	// 	fname: "amrit",
	// 	lname: "poudel",
	// 	age:   26,
	// }
	// p2 := person{
	// 	fname: "anita",
	// 	lname: "poude;",
	// 	age:   23,
	// }
	// fmt.Println(p1.fname, p1.lname, p1.age)
	// fmt.Println(p2)

	// we also can use anynomous struct for temporary purpose
	user := struct {
		fname string
		lname string
		age   int
	}{
		fname: "amrit",
		lname: "poudel",
		age:   20,
	}
	fmt.Println(user)

}
