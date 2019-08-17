package main

import (
	"encoding/json"
	"net/http"
	"errors"
	"fmt"
	"strings"
	"strconv"
)


// We define a Weapon struct to handle all rows of CSV
// - we define the fields with its type and how to parse
type Payment struct {
	Name             string  
	Amount           float64
	Lang             string
}


func CheckPaymentIsValid(payment Payment) error {
	if payment.Name == "" {
		errMsg := getCodeMessage("ERR_PAYMENT_NAME", payment.Lang)
		return errors.New(errMsg)
	}
	if payment.Amount <= 0 {
		errMsg := getCodeMessage("ERR_PAYMENT_AMOUNT", payment.Lang)
		return errors.New(errMsg)
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
func PerformPayment(payment Payment, walletAddress string) string{
	var msg string

	// If contact address is not defined, then finish the process
	contactAddress := GetContactAddress(payment.Name)
	if contactAddress == "" {
		// Generate a Human Response
		msg := getCodeMessage("ERR_NO_CONTACT_REGISTERED", payment.Lang)
		msg = strings.Replace(msg, "${0}", payment.Name, -1)
		return msg
	}
	
	// Check the balance of the account
	walletBalanceHex := GetAddressBalance(walletAddress)
	walletBalance := HexAmountWeiToEther(walletBalanceHex)
	hasEnoughFundsInWallet := walletBalance > payment.Amount
	if !hasEnoughFundsInWallet {
		msg := getCodeMessage("ERR_NO_ENOUGH_FUNDS", payment.Lang)
		msg = strings.Replace(msg, "${0}", strconv.FormatFloat(payment.Amount, 'f', 6, 64), -1)
		msg = strings.Replace(msg, "${1}", strconv.FormatFloat(walletBalance, 'f', 6, 64), -1)
		return msg
	}

	// Get nonce from account
	walletNonce, err := getAddressNonce(walletAddress)
	if err != nil {
		msg := getCodeMessage("ERR_OBTAIN_NONCE", payment.Lang)
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
		msg := getCodeMessage("ERR_SIGNING_TX", payment.Lang)
		return msg
	}
	fmt.Println("SIGNED TX: ", signedTx)
	

	// Send transaction to infura
	txHash, err := SendRawTransaction(signedTx)
	if err != nil {
		msg := getCodeMessage("ERR_SENDING_TX", payment.Lang)
		return msg
	}

	// All works fine! Tx done!
	fmt.Println("Hash from TX on Ropsten: ", txHash)
	msg = getCodeMessage("TX_SENT", payment.Lang)
	msg = strings.Replace(msg, "${0}", payment.Name, -1)
	msg = strings.Replace(msg, "${1}", fmt.Sprintf("%f", payment.Amount), -1)
	msg = strings.Replace(msg, "${2}", txHash, -1)
	return msg
}
