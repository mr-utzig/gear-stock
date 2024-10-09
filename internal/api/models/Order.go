package models

import "database/sql"

type Order struct {
	ID          int    `json:"id"`
	CustomerID  int    `json:"customer_id"`
	UserID      int    `json:"user_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	OrderDate   string `json:"order_date"`
	Status      bool   `json:"status"`
}

type OrderModel struct {
	db *sql.DB
}

func NewOrderModel(db *sql.DB) *OrderModel {
	return &OrderModel{db: db}
}

func (o *OrderModel) GetAllOrders() *[]Order {
	return &[]Order{}
}

func (o *OrderModel) CreateOrder() *Order {
	return &Order{}
}

func (o *OrderModel) GetOrder(id int) *Order {
	return &Order{}
}

func (o *OrderModel) UpdateOrder(id int) *Order {
	return &Order{}
}
