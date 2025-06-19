package recover

import "fmt"

func main() {
	fmt.Println("program started")
	process()
	fmt.Println("program ended")

}

func process() {
	fmt.Println("process has started executing")
	// lets handle panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered: ", r)
		}
	}()

	panic("something bad happened during execution")
	fmt.Println("process is finished")
}
