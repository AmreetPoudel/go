package formating

import "fmt"

func main() {
	// message := "my name is amrit poudel"
	// message2 := "i am \ntesting "
	// message3 := "i am \t testing"
	// message4 := `i will be \n \t unchanged`
	// fmt.Println(message)
	// fmt.Println("the length of string message is ", len(message))
	// for index, character := range message {
	// 	fmt.Printf("%d:%c\n", index, character)
	// }
	// fmt.Println(message2)
	// fmt.Println(message3)
	// fmt.Println(message4)
	// msg1 := "good morning "
	// name := "amrit"
	// msg := msg1 + name
	// fmt.Println(msg)

	////string compare
	// a := "apple"
	// b := "banana"
	// fmt.Println(a < b) //check based on lexical graphical ordering or dictnary ordering
	//// ascii value of A=65 and a is 97

	//error
	err := age(25)
	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println("drive")
	}

}

func age(num int) error {
	if num < 19 {
		return fmt.Errorf("age %d is too young to drive .", num)

	}
	return nil

}
