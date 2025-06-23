// tried to explain myself with example
// lets say only esewa exists in nepal lets integrate esewa to our code
// 2nd case now khalti came now i need to integrate khalti too
// oops too much mess
// if another integration come them boom again if else ladder

package interface2

import "fmt"

type esewa struct {
	token string
}

func (e esewa) pay(amount float64) {
	//payement logic goes here
	fmt.Printf("paying amount %f using esewa token %s ", amount, e.token)
}

type khalti struct {
	token string
}

func (k khalti) pay(amount float64) {
	// payment logic goes here
	fmt.Printf("paying amount %f using khalti token %s ", amount, k.token)
}

func main() {
	// payment mode  and amount  get from UI
	payment_mode := "khalti"
	amount := 1000.0
	if payment_mode == "esewa" {
		e1 := esewa{
			//this all data will be populated from db connect and fetch
			//here we are hardcoding
			token: "adsf-adfad-af23-sdfa-23cs-afddsf",
		}

		e1.pay(amount)
	} else if payment_mode == "khalti" {

		k1 := khalti{
			//this all data will be populated from db connect and fetch
			//here we are hardcoding
			token: "adsf-adfad-af23-sdfa-23cs-afddsf",
		}

		k1.pay(amount)

	} else {
		fmt.println("invalid payemtn option")
	}

}
