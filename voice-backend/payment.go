package main

import (
	"encoding/json"
	"net/http"
	"errors"
	"fmt"
)


// We define a Weapon struct to handle all rows of CSV
// - we define the fields with its type and how to parse
type Payment struct {
	Name             string  
	Amount           float32  
}


func CheckPaymentIsValid(payment Payment) error {
	if payment.Name == "" {
		return errors.New("Payment field 'name' is not defined")
	}
	if payment.Amount == 0 {
		return errors.New("Payment field 'amount' is empty")
	}
	return nil
}


// Obtains the payment parameters as Payment struct
func ObtainPaymentParameters(request *http.Request) Payment {
	// Decode the request
	decoder := json.NewDecoder(request.Body)
	// Read the payment and check if read is ok
	var payment Payment
	
	err := decoder.Decode(&payment)
	if err != nil {
		panic(err)
	}

	return payment
}


// Perform the payment againts INFURA
func PerformPayment(payment Payment) string{
	var msg string

	contactAddress := GetContactAddress(payment.Name)
	// If contact address is not defined, then finish the process
	if contactAddress == "" {
		// Generate a Human Response
		msg = "Contacto " + payment.Name + " no tiene una dirección registrada"
	} else {
		// Call to infura
		txHash := sendTransactionToInfura(payment)
		msg = payment.Name + " va a recibir " + fmt.Sprintf("%f", payment.Amount) + " ether en su wallet. Mira aquí la transacción: https://ropsten.etherscan.io/tx/" + txHash
	}
	return msg
}

func sendTransactionToInfura(payment Payment) string {
	// TODO sign transaction
	// 		Send to infura
	// 		return tx hash

	mockTxHash := "0xa223f1cef82127af1ad545793f22fc0f3200be78de84dd82428898ea3adab602"
	return mockTxHash
}