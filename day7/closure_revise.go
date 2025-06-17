package closure_revise

import "fmt"

func main() {

	sub := subtractor()
	fmt.Println("the value of i is : ", sub())
	fmt.Println("the value of i is : ", sub())
	fmt.Println("the value of i is : ", sub())

}

func subtractor() func() int {
	i := 0
	fmt.Println("The value of I is ", i)
	return func() int {
		i++
		return i

	}
}
