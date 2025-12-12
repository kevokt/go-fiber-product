package models

import "time"

type Product struct {
	ID          int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	TokoID      int64     `gorm:"foreignKey:TokoID" json:"toko_id"`
	Product     string    `gorm:"type:varchar(255);not null" json:"product"`
	Description string    `gorm:"type:text" json:"description"`
	Quantity    int       `gorm:"type:int;not null" json:"quantity"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Toko        *Toko     `gorm:"foreignKey:TokoID" json:"toko,omitempty"`
}
