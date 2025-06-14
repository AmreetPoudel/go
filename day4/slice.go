package slice_practice

import "fmt"

func main() {
	// arr := [5]int{1, 2, 3, 4, 5}
	// fmt.Println(len(arr))
	// fmt.Println(arr)
	// slice1 := arr[1:4]

	// fmt.Println(slice1)
	// slice1[1] = 20
	// slice1[2] = 40
	// fmt.Println(len(arr))
	// fmt.Println(arr)
	// slices are nothing but array that point to original array it holds pointer to array,length current and total capacity
	//define slice
	// amritarr := [10]int{54, 65, 12, 76, 89, 2, 3, 12, 34, 65}
	// fmt.Println("original array: ", amritarr)
	// amritslice := amritarr[2:6]
	// fmt.Println("original slice: ", amritslice)
	// //making new slice
	// newslice := make([]int, 5)
	// copy(newslice, amritslice)
	// fmt.Println("dynamically created slice: ", newslice)
	// //lets change in slice then see difference in original slice we changed item in slice but original data
	// //changed because slice was just name or reference to original array so changes in one result change in another.

	// amritslice[2] = 1
	// fmt.Println("original slice: ", amritslice)
	// fmt.Println("original array: ", amritarr)

	// practice question
	// 1)
	// 	Create a slice of 5 ints. Print len and cap.
	// Append two more values. Print updated len and cap.
	// Slice [1:4] from an array. Print and modify it. Show array impact.
	amritarr := [10]int{1, 4, 2, 6, 3, 56, 1, 5, 6, 1}
	slice1 := amritarr[0:5]
	fmt.Println("slice : ", slice1)
	fmt.Println("length: ", len(slice1))
	fmt.Println("cap: ", cap(slice1))

	slice1 = append(slice1, 2, 3)
	fmt.Println("slice : ", slice1)
	fmt.Println("length: ", len(slice1))
	fmt.Println("cap: ", cap(slice1))

}
