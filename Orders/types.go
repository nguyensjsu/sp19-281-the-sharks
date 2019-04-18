package main

import (
	"gopkg.in/mgo.v2/bson"
)

type Items struct {
	ItemId      string `bson:"itemId" json:"itemId"`
	ItemName    string `bson:"itemName" json:"itemName"`
	Price       string `bson:"price" json:"price"`
	Description string `bson:"description" json:"description"`
}
type BurgerOrder struct {
	OrderId       bson.ObjectId `json:"orderId" bson:"_id"`
	UserId        string        `json:"userId" bson:"userId"`
	OrderStatus   string        `json:"orderStatus" bson:"orderStatus"`
	Items         []Items       `json:"items" bson:"items"`
	TotalAmount   float32       `json:"totalAmount" bson:"totalAmount"`
	PaymentStatus string        `bson:"PaymentStatus" json:"PaymentStatus"`
}
