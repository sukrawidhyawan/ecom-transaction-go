package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	UserID    int       `json:"user_id" gorm:"primary_key"`
	ProductID int       `json:"product_id" gorm:"primary_key"`
	Quantity  int       `json:"quantity"`
	Price     float32   `json:"price"`
	OrderTime time.Time `json:"order_time" gorm:"primary_key"`
}

func (Order) TableName() string {
	return "orders"
}
