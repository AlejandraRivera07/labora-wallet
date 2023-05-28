package models

import "time"

type Log struct {
	ID         string    `json:"id"`
	Date       time.Time `json:"date"`
	Status     string    `json:"status"`
	CustomerId string    `json:"national_id"`
	Country    string    `json:"country"`
	Codigo     string    `json:"check_id"`
}
