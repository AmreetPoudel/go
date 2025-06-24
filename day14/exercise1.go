package ex1

import "fmt"

type PaymentGateway interface {
	pay(amount float64)
}

type Esewa struct {
	token string
}

func (e Esewa) pay(amount float64) {
	fmt.Println("amount: ", amount, "paid . Token: ", e.token)
}

type Khalti struct {
	apiKey string
}

func (k Khalti) pay(amount float64) {
	fmt.Println("amount: ", amount, "paid . Token: ", k.apiKey)
}

func payment(g PaymentGateway, amount float64) {
	g.pay(amount)
}

func main() {
	var pg PaymentGateway

	pg = Esewa{
		token: "Esewa1234",
	}
	payment(pg, 1000)
	pg = Khalti{
		apiKey: "khalti1234",
	}
	payment(pg, 5000)
}
