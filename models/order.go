package models

import (
	"fmt"
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

type Orders []*Order

func (Order) TableName() string {
	return "orders"
}

func ListAllOrder() []*Order {
	order := make([]*Order, 0)
	err := GetDb().Table(Order{}.TableName()).Find(&order).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return order
}

func ListUserOrders(userId int) []*Order {
	order := make([]*Order, 0)
	err := GetDb().Table(Order{}.TableName()).Where("user_id = ? ", userId).Find(&order).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return order
}

func ListProductOrders(productId int) []*Order {
	order := make([]*Order, 0)
	err := GetDb().Table(Order{}.TableName()).Where("product_id = ? ", productId).Find(&order).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return order
}

func (orders Orders) CreateBulk() bool {
	for _, o := range orders {
		o.OrderTime = time.Now()
		err := GetDb().Create(o)
		if err != nil {
			fmt.Println(err)
		}
	}

	return true
}
