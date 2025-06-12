package main

import (
	"fmt"
)

func main() {
	// 1.check number is even or odd
	// var number int
	// fmt.Print("enter number: ")
	// fmt.Scanln(&number)
	// if number%2 == 0 {
	// 	fmt.Println("number is even")
	// } else {
	// 	fmt.Println("number is odd")
	// }

	// 	2. Age Group Classifier
	// Ask the user to enter their age.
	// If age < 13 → Print "Child"
	// If age between 13–19 → Print "Teen"
	// If age 20–59 → Print "Adult"
	// Else → Print "Senior"

	// var age int
	// fmt.Print("enter your age: ")
	// fmt.Scanln(&age)
	// if age <= 13 {
	// 	fmt.Printf("your age is %d so you are child ", age)
	// } else if age > 13 && age <= 19 {
	// 	fmt.Printf("your age is %d so you are teen ", age)
	// } else if age > 19 && age <= 59 {
	// 	fmt.Printf("your age is %d so you are adult ", age)
	// } else {
	// 	fmt.Println("your age is %d so you are senior ", age)
	// }

	// 	3. Vowel or Consonant
	// Ask the user to input a single character.
	// Print "Vowel" if it's a vowel (a, e, i, o, u), else print "Consonant".
	// (Ignore case sensitivity.)

	// var character string
	// fmt.Print("enter the character you want to check: ")
	// fmt.Scanln(&character)
	// if len(character) != 1 {
	// 	fmt.Println("enter single character")
	// 	return
	// }
	// lower_char := strings.ToLower(character)
	// if lower_char == "a" || lower_char == "e" || lower_char == "i" || lower_char == "o" || lower_char == "u" {
	// 	fmt.Printf("your entered character %v is vowel\n", character)
	// } else {
	// 	fmt.Printf("your entered character %v is consonent\n", character)
	// }

	//case practice
	var day int
	fmt.Print("enter day of week in number: ")
	fmt.Scan(&day)
	switch day {
	case 1:
		fmt.Println("its sunday")
	case 2:
		fmt.Println("its monday")
	case 3:
		fmt.Println("its tuesday")
	case 4:
		fmt.Println("its wednesday")
	case 5:
		fmt.Println("its thursday")
	case 6:
		fmt.Println("its friday")
	case 7:
		fmt.Println("its saturday")
	default:
		fmt.Println("wrong input")
	}

}
