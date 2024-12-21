package models

import (
	"github.com/mr-utzig/gear-stock/internal/database"
)

type Customer struct {
	ID    int64  `json:"id"`
	Doc   string `json:"doc"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type CustomerCreateRequest struct {
	Doc   string `json:"doc"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type CustomerUpdateRequest struct {
	Doc   string `json:"doc"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type CustomerModel struct {
}

func NewCustomerModel() *CustomerModel {
	return new(CustomerModel)
}

func (c *CustomerModel) GetAllCustomers() ([]Customer, error) {
	rows, err := database.Turso.Query("SELECT customer_id, document, name, email, phone FROM customer")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	customers := []Customer{}
	for rows.Next() {
		customer := Customer{}
		err := rows.Scan(&customer.ID, &customer.Doc, &customer.Name, &customer.Email, &customer.Phone)
		if err != nil {
			return nil, err
		}

		customers = append(customers, customer)
	}

	return customers, nil
}

func (c *CustomerModel) CreateCustomer(data *CustomerCreateRequest) (*Customer, error) {
	stmt, err := database.Turso.Prepare("INSERT INTO user (document, name, email, phone) VALUES (?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	exec, err := stmt.Exec(data.Doc, data.Name, data.Email, data.Phone)
	if err != nil {
		return nil, err
	}

	id, err := exec.LastInsertId()
	if err != nil {
		return nil, err
	}

	customer := &Customer{
		ID:    id,
		Doc:   data.Doc,
		Name:  data.Name,
		Email: data.Email,
		Phone: data.Phone,
	}

	return customer, nil
}

func (c *CustomerModel) GetCustomer(id int) (*Customer, error) {
	customer := new(Customer)
	err := database.Turso.QueryRow(
		`SELECT customer_id, document, name, email, phone
		FROM customer
		WHERE customer_id=?`,
		id,
	).Scan(
		&customer.ID,
		&customer.Doc,
		&customer.Name,
		&customer.Email,
		&customer.Phone,
	)

	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (c *CustomerModel) UpdateCustomer(id int, data *CustomerUpdateRequest) (*Customer, error) {
	stmt, err := database.Turso.Prepare("UPDATE customer SET document=?, name=?, email=?, phone=? WHERE customer_id=?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(data.Doc, data.Name, data.Email, data.Phone, id)
	if err != nil {
		return nil, err
	}

	customer, err := c.GetCustomer(id)
	if err != nil {
		return nil, err
	}

	return customer, nil
}
