package main

import (
	"encoding/json"
	"net/http"
	"bytes"
	"io/ioutil"
	"errors"
	"fmt"
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
		hasNoError := responseObject.Error == InfuraErrorResponse{}
		return hasNoError
	}
}



func getAddressNonce(address string) (string, error) {
	// Create Tx Request with rawTx as parameter
	requestForInfura := InfuraRequest {
		JsonRpc: "2.0",
		Method: "eth_getTransactionCount",
		Params: []string{address, "pending"},
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
		hasError := responseObject.Error != InfuraErrorResponse{}
		if !hasError {
			return responseObject.Result, nil
		} else {
			return "", errors.New(responseObject.Error.Message)
		}
		
	}
}


func GetAddressBalance(address string) (string) {
	// Create Tx Request with rawTx as parameter
	requestForInfura := InfuraRequest {
		JsonRpc: "2.0",
		Method: "eth_getBalance",
		Params: []string{address, "pending"},
		Id: 1,
	}
	requestBody, err := json.Marshal(requestForInfura)
	if err != nil {
		return "0.0"
	}

	// Make request
	resp, err := http.Post(INFURA_URL, "application/json", bytes.NewBuffer(requestBody) )
	if err != nil {
		return "0.0"
	}

	// Read body response
	if resp.StatusCode != 200 {
		return "0.0"
	} else {
		responseData ,err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "0.0"
		}
		
		// Unmarshall  - If contains error message, return false (Node not ready)
		var responseObject InfuraResponse
		json.Unmarshal(responseData, &responseObject)

		// check if error is empty
		hasError := responseObject.Error != InfuraErrorResponse{}
		if !hasError {
			return responseObject.Result
		} else {
			return "0.0"
		}	
	}

}



/**
	Send a raw transaction to Infura. Retrieve the txHash as result or error.
*/
func SendRawTransaction(rawTx string) (string, error) {
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
		hasError := responseObject.Error != InfuraErrorResponse{}
		if !hasError {
			return responseObject.Result, nil
		} else {
			fmt.Println("** Error while sending tx to Ropsten network. Cause:", responseObject.Error.Message)
			return "", errors.New(responseObject.Error.Message)
		}
		
	}

}