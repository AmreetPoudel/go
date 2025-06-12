package while

import (
	"fmt"
)

func main() {

	x := 10
	for true {
		fmt.Println(x)
		if x > 100 {
			break
		}
		x += 10
	}

}
