package practice

import "fmt"

func main() {

	// 1.Create an array of 5 integers and print them using a for loop.

	// 	fmt.Println("practicing")
	// 	amritarr := [5]int{}

	// 	for i := 0; i < 5; i++ {
	// 		fmt.Printf("item %d: ", i)
	// 		fmt.Scan(&amritarr[i])

	// 	}
	// 	for i := 0; i < 5; i++ {
	// 		fmt.Printf("item %d : %d\n", i, amritarr[i])
	// 	}
	// 2. Reverse an array using a loop and print the reversed values.

	// fmt.Println("printing array in revese order")
	// // define array
	// amritarr := [5]int{23, 13, 65, 76, 1}
	// var reversearr [5]int
	// for i := 0; i < len(amritarr); i++ {
	// 	reversearr[len(amritarr)-i-1] = amritarr[i]
	// }
	// fmt.Print(reversearr)

	//3.Find the sum of all elements in an integer array.
	// amritarr := [5]int{12, 56, 12, 17, 91}
	// var sum int
	// for i := range len(amritarr) {
	// 	sum = sum + amritarr[i]
	// }
	// fmt.Println("sum= ", sum)

	// Find the maximum value in an array of integers.
	// amritarr := [5]int{12, 56, 92, 17, 91}
	// var max_item = amritarr[0]
	// for i := range len(amritarr) {
	// 	if amritarr[i] > max_item {
	// 		max_item = amritarr[i]
	// 	}
	// }
	// fmt.Printf("max item is %d\n", max_item)
	// _ variable

	// amritarr := [10]int{12, 64, 243, 43, 9, 7934, 23, 345, 67, 87}
	// var total_divisible_by_3 int
	// for _, item := range amritarr {
	// 	if item%3 == 0 {
	// 		total_divisible_by_3 += 1
	// 	}
	// }
	// fmt.Printf("total number of items divisible by three are %d\n", total_divisible_by_3)
	// nums := [10]int{7, 14, 21, 28, 35, 42, 49, 56, 63, 70}
	// var number_greater_30 int

	// for _, num := range nums {

	// 	if num%2 == 0 {
	// 		fmt.Println(num)
	// 		if num > 30 {
	// 			number_greater_30 += 1
	// 		}
	// 	}
	// }
	// fmt.Println("total number greate then 30 and in even part: ", number_greater_30)

	nums := [10]int{5, 14, 67, 2, 99, 1, 33, 80, 42, 31}
	// here it is explicitely said i just need to count number of even index greater then 30 and print all item of even index
	var count int = 0
	var greater int = 0
	for _, num := range nums {
		if count%2 == 0 {
			fmt.Println(num)
			if num > 30 {
				greater += 1
			}
		}
		count += 1

	}
	fmt.Println("total greater number are: ", greater)

}
