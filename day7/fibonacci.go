package main

import "fmt"

func main() {
	// num := fibonacci(19)
	// fmt.Println(num)
	// this is recursion example now i will do normally then also print everything in order
	var total int
	first_number := 0
	second_number := 1
	fmt.Print("Enter total numbers up to which you want to print")
	fmt.Scan(&total)

	for i := 0; i < total; i++ {
		var new_number int
		print(first_number, " ")
		new_number = first_number + second_number
		first_number = second_number
		second_number = new_number
	}
}

// func fibonacci(number int) int {
// 	if number < 2 {
// 		return number
// 	} else {
// 		return fibonacci(number-1) + fibonacci(number-2)
// 	}
// }
