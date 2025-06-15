package func2

import "fmt"

func main() {
	// get two numbers from user
	var a, b int
	fmt.Print("enter first number: ")
	fmt.Scan(&a)
	fmt.Print("enter second number: ")
	fmt.Scan(&b)

	//sum
	fmt.Println("The sum of number is: ", operate(a, b, add))
	fmt.Println("the difference of number is: ", operate(a, b, sub))
	fmt.Println("the product of numbers is: ", operate(a, b, prod))
	fmt.Println("the div of numbers is: ", operate(a, b, div))

}

func operate(a int, b int, xxx func(int, int) int) int {
	return xxx(a, b)
}

func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func prod(a int, b int) int {
	return a * b
}
func div(a int, b int) int {
	return a / b
}
