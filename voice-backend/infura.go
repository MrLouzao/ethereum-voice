package main

import (
	"encoding/json"
	"net/http"
	"bytes"
	"io/ioutil"
	"errors"
)


// Typical infura request
type InfuraRequest struct {
	JsonRpc		string
	Method 		string
	Params 		[]string
	Id 			int
}
// OK Response from Infura
type InfuraResponse struct {
	JsonRpc 	string					`json:"jsonrpc"`
	Id			int						`json:"id"`
	Result		string					`json:"result"`
	Error 		InfuraErrorResponse		`json:"error"`
}
// KO Response from Infura
type InfuraErrorResponse struct {
	Id			int		`json:"id"`
	Message		string	`json:"message"`
}


func checkInfuraStatus() bool {
	// Build object to send
	requestForInfura := InfuraRequest {
		JsonRpc: "2.0",
		Method: "eth_chainId",
		Params: []string{},
		Id: 1,
	}
	requestBody, err := json.Marshal(requestForInfura)
	if err != nil {
		return false
	}

	// Make request
	resp, err := http.Post(INFURA_URL, "application/json", bytes.NewBuffer(requestBody) )
	if err != nil {
		return false
	}

	// Read body response
	if resp.StatusCode != 200 {
		return false
	} else {

		responseData ,err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return false
		}

		// Unmarshall  - If contains error message, return false (Node not ready)
		var responseObject InfuraResponse
		json.Unmarshal(responseData, &responseObject)
		// check if error is empty
		result := responseObject.Error == InfuraErrorResponse{}
		return result
	}
}



/**
	Send a raw transaction to Infura. Retrieve the txHash as result or error.
*/
func sendRawTransaction(rawTx string) (string, error) {
	// Create Tx Request with rawTx as parameter
	requestForInfura := InfuraRequest {
		JsonRpc: "2.0",
		Method: "eth_sendRawTransaction",
		Params: []string{rawTx},
		Id: 1,
	}
	requestBody, err := json.Marshal(requestForInfura)
	if err != nil {
		return "", errors.New("Failing at creating request")
	}

	// Make request
	resp, err := http.Post(INFURA_URL, "application/json", bytes.NewBuffer(requestBody) )
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
		var responseObject InfuraResponse
		json.Unmarshal(responseData, &responseObject)
		// check if error is empty
		hasError := responseObject.Error == InfuraErrorResponse{}
		if !hasError {
			return responseObject.Result, nil
		} else {
			return "", errors.New(responseObject.Error.Message)
		}
		
	}

	return "Everything is ok", nil
}