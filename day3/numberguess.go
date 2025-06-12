package game

import (
	"fmt"
	"math/rand"
)

func main() {

	fmt.Println("***welcome to number guessing game***")
	fmt.Println("system is generating random number..")
	random_number := rand.Intn(100)
	fmt.Println("system generated random number..")
	// fmt.Println(random_number)
	total_attempts := 0
	var guess_number int
	for {

		if total_attempts == 5 {
			fmt.Println("your total change is over")
			fmt.Println("correct number was ", random_number)
			break
		}

		fmt.Println("-----------------------------------")
		fmt.Println("total guess remaining: ", 5-total_attempts)
		fmt.Print("guess random number: ")
		_, err := fmt.Scanln(&guess_number)
		if err != nil {
			fmt.Println("enter number not string")
			continue
		}

		if guess_number > random_number {
			fmt.Println("guess smaller number")
		}
		if guess_number < random_number {
			fmt.Println("guess large number")
		}
		if guess_number == random_number {
			fmt.Println("good guess ")
			fmt.Printf("random number was: %d ", random_number)
			break

		}
		total_attempts += 1

	}
}
