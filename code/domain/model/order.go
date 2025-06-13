package model

import "time"

type Order struct {
	ID         string    `bson:"_id"`
	CustomerID string    `bson:"customer_id"`
	Total      float64   `bson:"total"`
	CreatedAt  time.Time `bson:"created_at"`
}

func NewOrder(id, customerID string, total float64) *Order {
	return &Order{
		ID:         id,
		CustomerID: customerID,
		Total:      total,
		CreatedAt:  time.Now(),
	}
}
