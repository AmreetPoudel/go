package interface4

import "fmt"

// create interface to use
// here pay is something all payment gateway do so we crete pay method for interface
type PaymentGateway interface {
	pay(amount float64)
}

// create a struct and function to consume this interface
// wallet1 esewa
type esewa struct {
	token string
}

func (e esewa) pay(amount float64) {
	fmt.Println("paying amount: ", amount, "using esewa. Key :", e.token)
}

// similar for khalti
type khalti struct {
	apiKey string
}

func (k khalti) pay(amount float64) {
	fmt.Println("paying amount: ", amount, "using khalti. Key:", k.apiKey)
}

// 3. common function that works with any payment gateway
func ProcessPayment(pg PaymentGateway, amount float64) {
	pg.pay(amount)
}

func main() {

	//4. declearing variable of interface type

	var pg PaymentGateway

	// You can assign any struct that has Pay() method
	pg = esewa{
		token: "ESEWA123",
	}
	ProcessPayment(pg, 2000)

	pg = khalti{
		apiKey: "KHALTI1234",
	}
	ProcessPayment(pg, 4000)

}
