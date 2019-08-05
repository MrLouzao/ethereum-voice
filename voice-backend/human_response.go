package main

// Human response for Google Assistant
// - If we want to marshall the field fulfillment on a json response,
//		we must put the struct field name in capital letter and we must place the json property at the end of the property
type HumanResponse struct {
	FulfillmentText string `json:"fulfillmentText"`
}