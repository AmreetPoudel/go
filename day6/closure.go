package closure

import "fmt"

func main() {
	//lets assign the returned function to variable
	sequence := adder()
	fmt.Println("1.the value of i is : ", sequence())
	fmt.Println("1.the value of i is : ", sequence())
	fmt.Println("1.the value of i is : ", sequence())

}

func adder() func() int {
	i := 0
	fmt.Println("the value of i is : ", i)
	return func() int {
		i++
		fmt.Println("added 1 to i")
		return i
	}
}

//explanatin i am not 100% SURE What i am doing but
//here we returned a anynomous function but then value i is defined outside the anynomous functin
// now we are putting that anynomous function to variable
// that anynomous functins simple increamtnt value of i but it is not its value
//when we call sequence()  ,adder() has returned anynomous functon so now sequence is anynomus functin that is retuned
// here i++ (returned) it is pointing to same memory address  so sequence is getting function and its environment
//too that is closure
//That function + captured variable = closure.
