package slice

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(len(arr))
	fmt.Println(arr)
	slice1 := arr[1:4]

	fmt.Println(slice1)
	slice1[1] = 20
	slice1[2] = 40
	fmt.Println(len(arr))
	fmt.Println(arr)

}
