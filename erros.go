package error1

import (
	"errors"
	"fmt"
	"math"
)

func SquareRoot(number float64) (float64, error) {
	if number < 0 {
		return 0, errors.New("math error: cannot compute squareroot of negative number")
	} else {
		return math.Sqrt(number), nil
	}

}

func main() {
	number := 11.11
	n, e := SquareRoot(float64(number))

	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Printf("The square root of given number: %v is %.2f\n", number, n)
	}

}
