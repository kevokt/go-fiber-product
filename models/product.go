package models

import "time"

type Product struct {
	ID          int64     `json:"id"`
	Product     string    `json:"product"`
	Description string    `json:"description"`
	Quantity    int       `json:"quantity"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
