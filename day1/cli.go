package main

import "fmt"

func main() {
	var name string
	var age int
	var height float64
	var is_learning_golang bool

	fmt.Println("Enter you details below")
	fmt.Print("Name: ")
	fmt.Scan(&name)
	fmt.Print("age: ")
	fmt.Scan(&age)
	fmt.Print("height: ")
	fmt.Scan(&height)
	fmt.Print("are you learning golang: ")
	fmt.Scan(&is_learning_golang)
	fmt.Println("printing your details..")
	fmt.Println("name: ", name)
	fmt.Println("age: ", age)
	fmt.Println("height: ", height)
	fmt.Println("learning golang: ", is_learning_golang)

}
