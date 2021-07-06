package dto

import "time"

// Layer that recieves externaldata, later they will be imported in a usecase
type Transaction struct {
	ID              string
	Name            string
	Number          string
	ExpirationMonth int32
	ExpirationYear  int32
	CVV             int32
	Amount          float64
	Store           string
	Description     string
	CreatedAt       time.Time
}