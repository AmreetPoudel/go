package error2

import "fmt"

func main() {
	err := process()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("task completed successfully")
}

type mycustomerror struct {
	status_code int
	message     string
}

// this will return Error . we are implementing Error() of error interface
func (err *mycustomerror) Error() string {
	return fmt.Sprintf("oops!! you got an error. \n status_code:%d message: %s", err.status_code, err.message)
}

func process() error {
	return &mycustomerror{
		status_code: 404,
		message:     "not found",
	}
}
