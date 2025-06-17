package fact1

import "fmt"

func main() {
	fact := factorial(5)

	fmt.Println(fact)

	sum := sum_of_digits((192))
	fmt.Println(sum)
}

func factorial(num int) int {
	if num == 0 || num == 1 {
		return 1
	} else {
		return num * factorial(num-1)
	}
}

// sum of digits of any number
func sum_of_digits(num int) int {
	if num < 10 {
		return num
	} else {
		return num%10 + sum_of_digits(num/10)
	}
}
