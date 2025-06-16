package panic

import "fmt"

func main() {
	process(-5)
}

func process(number int) {
	fmt.Println("process is running")
	defer fmt.Println("no matter i will run at last but before panic this time")
	defer fmt.Println("i will run at 2nd last and before panic also")
	if number < 0 {
		fmt.Println("panic starting")
		panic("cannot operate on negative number")

	} else {
		fmt.Println("processing number :", number)
	}

}
