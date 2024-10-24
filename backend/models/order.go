package models

type Order struct {
	ID          int 	`json:"id"`
	Customer    string	`json:"customer"`
	Amount      float64	`json:"amount"`
	OrderStatus string	`json:"orderStatus"`

}