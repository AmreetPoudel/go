package reverse

import "fmt"

func main() {
	// rev := reverse_int(12345)
	// fmt.Println(rev)

	var n int = 12345
	sum := 0
	for n != 0 {
		rem := n % 10
		sum = sum*10 + rem
		n = n / 10
	}
	fmt.Println("the reverse number is :", sum)

}

// func reverse_int(n int) int {
// 	if n == 0 {
// 		return 0
// 	} else {
// 		return n%10 + reverse_int(n/10)
// 	}
// }

//this is wrong  so commented did this problem via simple logic
