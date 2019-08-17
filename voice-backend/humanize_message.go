package main


// Obtain the message correspondant to the code in the proper language. English by default.
func getCodeMessage(code string, lang string) string {
	if lang == "es" {
		return esMessages[code]
	} else {
		// lang == "en"
		return enMessages[code]
	}
}