package main

import (
	"encoding/json"
	"net/http"
	"io/ioutil"
	"errors"
	"fmt"
)


// Represents an unsigned transaction
type EthereumTransaction struct {
	Nonce		string
	Recipient	string
	Value		string
}


type SignedTransactionResponse struct {
	Status		int
	Data		string
	Message		string
}


/**
	Sign a transaction on signer backend
*/
func SignRawTransaction(txToSign EthereumTransaction) (string, error) {

	payloadParameters := fmt.Sprintf("/signedTransaction?transaction={\"nonce\":\"%s\",\"gasPrice\":\"0xb2d05e00\",\"gasLimit\":\"0x27100\",\"to\":\"%s\",\"value\":\"%s\",\"data\":\"0x00\"}", txToSign.Nonce, txToSign.Recipient, txToSign.Value)
	signerUrlWithData := SIGNER_URL + payloadParameters
	fmt.Println("URL TO SIGN: ", signerUrlWithData)


	// Make request
	resp, err := http.Get(signerUrlWithData)
	if err != nil {
		return "", errors.New("Failing at sending request to Infura")
	}

	// Read body response
	if resp.StatusCode != 200 {
		return "", errors.New("Failing at sending request to Infura")
	} else {
	
		responseData ,err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", errors.New("Failing at parsing response from Infura")
		}
		
		// Unmarshall  - If contains error message, return false (Node not ready)
		var responseObject SignedTransactionResponse
		json.Unmarshal(responseData, &responseObject)

		// check if error is empty
		return responseObject.Data, nil
	
	}
}