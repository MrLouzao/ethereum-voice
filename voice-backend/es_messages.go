package main

var esMessages = map[string]string{
	"ERR_PAYMENT_NAME": "El pago no reconoce el recipiende",
	"ERR_PAYMENT_AMOUNT": "El pago no reconoce la cantidad a enviar",
	"ERR_NO_CONTACT_REGISTERED": "El contacto ${0} no tiene una dirección registrada",
	"ERR_NO_ENOUGH_FUNDS": "No hay fondos suficientes en la wallet. Requeridos ${0} pero actualmente solo disponibles ${1}",
	"ERR_OBTAIN_NONCE": "Se ha producido un error al obtener los parámetros de tu cuenta",
	"ERR_SIGNING_TX": "Se ha producido un error al firmar la transacción. La transacción no se efectuará.",
	"ERR_SENDING_TX": "Se ha producido un error al enviar la transacción a la red. La transacción no se efectuará.",
	"TX_SENT": "El destinatario ${0} va a recibir ${1} ether en su wallet. Mira aquí el recibo de la transacción: \n https://ropsten.etherscan.io/tx/${2}",
}