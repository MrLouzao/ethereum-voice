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
	Amount           float64  
}


func CheckPaymentIsValid(payment Payment) error {
	if payment.Name == "" {
		return errors.New("Payment field 'name' is not defined")
	}
	if payment.Amount <= 0 {
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
// TODO
// 1. Get the recipient
// 2. Check the balance for from account
// 3. Get the nonce for address
// 4. Create tx
// 5. Sign tx calling the signer
// 6. Send tx to Infura
// 7. Return tx hash from network
func PerformPayment(payment Payment, walletAddress string) string{
	var msg string

	// If contact address is not defined, then finish the process
	contactAddress := GetContactAddress(payment.Name)
	if contactAddress == "" {
		// Generate a Human Response
		msg := "Contacto " + payment.Name + " no tiene una dirección registrada"
		return msg
	}
	
	// Check the balance of the account
	walletBalanceHex := GetAddressBalance(walletAddress)
	walletBalance := HexAmountWeiToEther(walletBalanceHex)
	hasEnoughFundsInWallet := walletBalance > payment.Amount
	if !hasEnoughFundsInWallet {
		msg := fmt.Sprintf("No hay fondos suficientes en la wallet. Requeridos: %f pero actualmente en cuenta %f", payment.Amount, walletBalance)
		return msg
	}

	// Get nonce from account
	walletNonce, err := getAddressNonce(walletAddress)
	if err != nil {
		msg := "Se ha producido un error al obtener los parámetros de tu cuenta"
		return msg
	}


	// Sign the transaction
	amountInHex := EtherFloatToWeiHex(payment.Amount)
	rawTransaction := EthereumTransaction {
		Nonce: walletNonce,
		Recipient: contactAddress,
		Value:	amountInHex,
	}
	signedTx, err := SignRawTransaction(rawTransaction)
	if err != nil {
		msg := "Se ha producido un error al firmar la transacción"
		return msg
	}
	fmt.Println("SIGNED TX: ", signedTx)
	

	// Send transaction to infura
	txHash, err := SendRawTransaction(signedTx)
	if err != nil {
		msg := "Se ha producido un error al enviar la transacción a la red"
		return msg
	}

	// All works fine! Tx done!
	fmt.Println("EL HASH RECIBIDO ES: ", txHash)
	msg = payment.Name + " va a recibir " + fmt.Sprintf("%f", payment.Amount) + " ether en su wallet. Mira aquí la transacción: https://ropsten.etherscan.io/tx/" + txHash
	return msg
}
