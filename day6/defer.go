package deferexample

import "fmt"

func main() {
	//defer means postpond to last for example  a process is runing defer sstatement are put on stack and popped
	// out at last before function exits
	// defer functions  are evaulated immediately but only run at last so if we assign any value and print in differ
	// and that value changes then also previous value get printed as defer is evaluated sequentially at runtime
	// but performed at last
	process()

}

func process() {
	defer fmt.Println("differ task need to do at last")
	defer fmt.Println("i will print at second last")
	defer fmt.Println("i wll print at 3rd last")

	//variable example
	// fefer func is evaluated normally during execution only printed or task is done at lastls

	i := 100
	defer fmt.Println("value in defer: ", i)
	fmt.Println("i come first")
	i++
	fmt.Println("value normally", i)
}
