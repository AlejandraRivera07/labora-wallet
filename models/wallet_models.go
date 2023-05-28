package models

import "time"

// Estructura para representar una billetera
type Wallet struct {
	ID         string    `json:"id"`
	CustomerId string    `json:"personId" validate:"required"`
	CreateDate time.Time `json:"date" validate:"required"`
	CountryId  string    `json:"country"`
}
