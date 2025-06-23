package interface

import "fmt"

// =============================
// STEP 1: CREATE INTERFACE
// =============================

// Interface is like a "contract" or a "rule".
// Any struct that wants to be called a PaymentGateway
// must have a method named 'pay' with this exact signature:
// --> func pay(amount float64)
type PaymentGateway interface {
	pay(amount float64)
}

// =============================
// STEP 2: CREATE STRUCTS THAT IMPLEMENT THIS INTERFACE
// =============================

// Wallet 1: esewa
// This struct represents eSewa payment system.
// It has a token (like an API key).
type esewa struct {
	token string
}

// This is the 'pay' function for esewa.
// Since it matches the interface (same name + same parameter),
// Go will automatically treat 'esewa' as implementing the interface.
func (e esewa) pay(amount float64) {
	fmt.Println("paying amount:", amount, "using eSewa. Key:", e.token)
}

// Wallet 2: khalti
// Another struct, this time for Khalti payment system.
type khalti struct {
	apiKey string
}

// This is the 'pay' function for khalti.
// It also matches the interface method exactly,
// so Go knows khalti satisfies the PaymentGateway interface.
func (k khalti) pay(amount float64) {
	fmt.Println("paying amount:", amount, "using Khalti. Key:", k.apiKey)
}

// =============================
// STEP 3: COMMON FUNCTION USING THE INTERFACE
// =============================

// This function accepts anything that satisfies the PaymentGateway interface.
// That means it can take 'esewa', 'khalti', or any other struct
// as long as they have a pay(amount float64) function.
func ProcessPayment(pg PaymentGateway, amount float64) {
	// Call the pay method defined by whichever struct is passed.
	pg.pay(amount)
}

// =============================
// STEP 4: MAIN FUNCTION
// =============================

func main() {
	// Declare a variable of type 'PaymentGateway'.
	// This variable can hold any struct that has a 'pay(float64)' method.
	var pg PaymentGateway

	// Assign an esewa struct to the interface variable
	// Go checks: "Does esewa have a method named pay(amount float64)?"
	// Answer: YES → So this works.
	pg = esewa{
		token: "ESEWA123",
	}
	ProcessPayment(pg, 2000)

	// Now assign a khalti struct to the same interface variable.
	// Again Go checks the same rule, and since khalti also has
	// a matching pay() method → It also works.
	pg = khalti{
		apiKey: "KHALTI1234",
	}
	ProcessPayment(pg, 4000)
}
