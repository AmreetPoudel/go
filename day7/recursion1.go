package recursion1

import "fmt"

func main() {
	zeros := return_zeros(100100101)
	fmt.Println(zeros)

}

func return_zeros(n int) int {
	if n == 0 {
		return 0
	}
	if n%10 == 0 {
		return 1 + return_zeros(n/10)
	} else {
		return return_zeros(n / 10)
	}
}
