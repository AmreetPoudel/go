package test

import "fmt"

func main() {
	var numbers = [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("index: %d  value: %d \n", i, numbers[i])
	}

	for index, data := range numbers {
		fmt.Println(index, " -> ", data)
	}
	fmt.Println(numbers)
	var emptynumber [5]int
	fmt.Println(emptynumber)
	emptynumber[3] = 20
	fmt.Println(emptynumber)

}
