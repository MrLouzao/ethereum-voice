package main

// We must perform $go get for mux and text packages
import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"os"
	"fmt"
	"errors"
)

// Global vars
var INFURA_URL string
var SIGNER_URL string
var PORT string


// Handle the request of payment to client
func MakeEthereumPayment(w http.ResponseWriter, request *http.Request) {
	payment := ObtainPaymentParameters(request)

	err := CheckPaymentIsValid(payment)
	if err != nil {
		myHumanResponse := HumanResponse{FulfillmentText: err.Error()}
		json.NewEncoder(w).Encode(myHumanResponse)
	} else {
		// Perform the payment and send response to user
		stateMessage := PerformPayment(payment, "0x007BfF585Be4B690db8FBC9FDe4F936294Fa37De")
		myHumanResponse := HumanResponse{FulfillmentText: stateMessage}
		json.NewEncoder(w).Encode(myHumanResponse)
	}
}


func readEnvVars() (string, string, string, error){
	// port (8000 by default)
	var port string
	if os.Getenv("PORT") != "" {
		port = ":" + os.Getenv("PORT")
	} else {
		port = ":8000"
	}

	// infura URL
	var infuraUrl string
	if os.Getenv("INFURA_URL") != "" {
		infuraUrl = os.Getenv("INFURA_URL")
	} else {
		return "", "", "", errors.New("INFURA_URL not defined")
	}

	// signer URL
	var signerUrl string
	if os.Getenv("SIGNER_URL") != "" {
		signerUrl = os.Getenv("SIGNER_URL")
	} else {
		return "", "", "", errors.New("SIGNER_URL not defined")
	}

	return port, infuraUrl, signerUrl, nil
}


/**
mux creates a router at the very beginning
- We add the /weapons endpoint URL on GET method. GetWeapons is the function responsible of handling the request
- With :8000 we serve the application on port 8000
*/
func main() {
	// port (8000 by default)
	var err error
	PORT, INFURA_URL, SIGNER_URL, err = readEnvVars()
	if err != nil {
		panic(err)
	}
	fmt.Println("INFURA_URL=", INFURA_URL)
	fmt.Println("SIGNER_URL=", SIGNER_URL)
	fmt.Println("PORT=", PORT)

	// Check infura is up
	isInfuraUp := checkInfuraStatus()
	fmt.Println("** Is infura up?", isInfuraUp)

	/* // Print gas price in hex
	gasPrice := EtherFloatToWeiHex(0.000000003)
	fmt.Println("Current gas price: ", gasPrice)
	*/

	// Initialize
	router := mux.NewRouter()
	router.HandleFunc("/voice-payment", MakeEthereumPayment).Methods("POST")
	http.ListenAndServe(PORT, router)
}