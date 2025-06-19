package pointer

import "fmt"

func main() {
	a := 20
	var ptr *int
	var ptr1 *int
	ptr = &a
	fmt.Println("the pointer pointing value is ", *ptr)
	fmt.Println("actual pointer valueof address it pointting to is ", ptr)
	fmt.Println("address of a ", &a)
	increment(ptr)
	fmt.Println(a)
	fmt.Println("here we are seeing what empty pointer hold ")
	if ptr1 == nil {
		fmt.Println("it is pointing nill")
	}
}

func increment(temp *int) {
	*temp++
}
