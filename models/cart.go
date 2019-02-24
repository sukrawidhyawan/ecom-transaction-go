package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Cart struct {
	gorm.Model
	UserID    int `json:"user_id" gorm:"primary_key"`
	ProductID int `json:"product_id" gorm:"primary_key"`
	Quantity  int `json:"quantity"`
}

func (Cart) TableName() string {
	return "carts"
}

func ListAllCart() []*Cart {
	cart := make([]*Cart, 0)
	err := GetDb().Table(Cart{}.TableName()).Find(&cart).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return cart
}

func GetCart(userID int, productID int) *Cart {
	cart := &Cart{}
	err := GetDb().Table(Cart{}.TableName()).Where("user_id=? AND product_id=?", userID, productID).Find(&cart).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return cart
}

func (cart *Cart) Create() *Cart {
	err := GetDb().Create(cart)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return cart
}
