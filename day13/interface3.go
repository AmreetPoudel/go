// implement interface here

package main

import (
	"fmt"
)

type PaymentGateway interface {
	Pay(amount float64)
}

// esewa
type esewa struct {
	token string
}

func (e esewa) Pay(amount float64) {
	fmt.Printf("total amount paid : %f from esewa toekn: %s", amount, e.token)
}

// khalti
type khalti struct {
	token string
}

func (k khalti) Pay(amount float64) {
	fmt.Printf("total amount paid : %f from khalti toekn: %s", amount, k.token)
}

func ProcessPayment(pg PaymentGateway, amount float64) {
	pg.Pay(amount)
}

func main() {
	var pg PaymentGateway

	method := "khalti" // try changing to "khalti"

	if method == "esewa" {
		pg = esewa{token: "ESEWA123"}
	} else if method == "khalti" {
		pg = khalti{token: "KHALTI456"}
	}

	ProcessPayment(pg, 1500)
}
