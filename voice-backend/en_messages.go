package main

var enMessages = map[string]string{
	"ERR_PAYMENT_NAME": "The payment does not recognize the recipient",
	"ERR_PAYMENT_AMOUNT": "The payment does not recognize the amount",
	"ERR_NO_CONTACT_REGISTERED": "Contact ${0} has no address registered",
	"ERR_NO_ENOUGH_FUNDS": "Not enough funds in your wallet. Required ${0} but only ${1} available",
	"ERR_OBTAIN_NONCE": "Error while obtaining the parameters of your account",
	"ERR_SIGNING_TX": "An error occurred when trying to sign the transaction. The transaction will not take place.",
	"ERR_SENDING_TX": "An error occurred when trying to send the transaction to the network. The transaction will not take place",
	"TX_SENT": "The recipient ${0} is going to receive ${1} ether in it's wallet. Check here the transaction receipt: \n https://ropsten.etherscan.io/tx/${2}",
}