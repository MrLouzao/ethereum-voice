package main

// We must perform $go get for mux and text packages
import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"os"
)


// Handle the request of payment to client
func MakeEthereumPayment(w http.ResponseWriter, request *http.Request) {
	payment := ObtainPaymentParameters(request)

	err := CheckPaymentIsValid(payment)
	if err != nil {
		myHumanResponse := HumanResponse{FulfillmentText: err.Error()}
		json.NewEncoder(w).Encode(myHumanResponse)
	} else {
		// Perform the payment and send response to user
		stateMessage := PerformPayment(payment)
		myHumanResponse := HumanResponse{FulfillmentText: stateMessage}
		json.NewEncoder(w).Encode(myHumanResponse)
	}
}


/**
mux creates a router at the very beginning
- We add the /weapons endpoint URL on GET method. GetWeapons is the function responsible of handling the request
- With :8000 we serve the application on port 8000
*/
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/voice-payment", MakeEthereumPayment).Methods("POST")

	// port (8000 by default)
	var port string

	os.Getenv("PORT")
	if os.Getenv("PORT") != "" {
		port = ":" + os.Getenv("PORT")
	} else {
		port = ":8000"
	}
	http.ListenAndServe(port, router)
}