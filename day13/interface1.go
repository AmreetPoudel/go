// tried to explain myself with example
// lets say only esewa exists in nepal lets integrate esewa to our code

package test

import "fmt"

type esewa struct {
	token string
}

func (e esewa) pay(amount float32) {
	//payement logic goes here
	fmt.Printf("paying amount %f using esewa token %s ", amount, e.token)
}

func main() {

	e1 := esewa{
		//this all data will be populated from db connect and fetch
		//here we are hardcoding
		token: "adsf-adfad-af23-sdfa-23cs-afddsf",
	}

	e1.pay(1000)

}
