package loop

func main() {

	// // simple for loop

	// for i := 0; i <= 5; i++ {
	// 	fmt.Println(i)
	// }

	// // looping over list of items
	// numbers := []int{1, 2, 3, 4, 5}
	// for index, value := range numbers {
	// 	fmt.Printf("index: %d  value: %d\n", index, value)
	// }

	// //printing only odd numbers
	// for i := 0; i <= 10; i++ {
	// 	if i%2 == 0 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	// // break if encountered number 7
	// fmt.Println("printing number but break if 7 encountered")
	// for i := 0; i <= 10; i++ {
	// 	if i == 7 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }

	// printing pattern
	//     *
	//    ***
	//   *****
	//  *******
	// *********

	// rows := 5
	// for i := 1; i <= rows; i++ {
	// 	// print spaces
	// 	for j := i; j <= rows; j++ {
	// 		print(" ")
	// 	}
	// 	//print *s
	// 	for k := 1; k <= 2*i-1; k++ {
	// 		print("*")
	// 	}
	// 	fmt.Println()
	// }

	// another pattern

	//  ********* 9 stars 1 space
	//   *******  7 stars	2 spaces
	//    *****		5 stars	3spaces
	//     ***		3 stars	4spaces
	//      *		1 star	5 spaces
	//rows=5

	for i := range 5 {
		// print spaces
		for j := 1; j <= i; j++ {
			print(" ")
		}
		// print *s

		for k := 1; k < 2*(5-i); k++ {
			print("*")
		}
		println()
	}
	// for i := range 10 {
	// 	fmt.Println(i)
	// }

}
