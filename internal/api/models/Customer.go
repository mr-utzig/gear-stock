package models

import "database/sql"

type Customer struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type CustomerModel struct {
	db *sql.DB
}

func NewCustomerModel(db *sql.DB) *CustomerModel {
	return &CustomerModel{db: db}
}

func (o *CustomerModel) GetAllCustomers() *[]Customer {
	return &[]Customer{}
}

func (o *CustomerModel) CreateCustomer() *Customer {
	return &Customer{}
}

func (o *CustomerModel) GetCustomer(id int) *Customer {
	return &Customer{}
}

func (o *CustomerModel) UpdateCustomer(id int) *Customer {
	return &Customer{}
}
