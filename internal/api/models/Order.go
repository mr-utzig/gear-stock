package models

import "github.com/mr-utzig/gear-stock/internal/database"

type Order struct {
	ID          int64  `json:"id"`
	CustomerID  int    `json:"customer_id"`
	UserID      int    `json:"user_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	OrderDate   string `json:"order_date"`
	Status      bool   `json:"status"`
}

type OrderCreateRequest struct {
	CustomerID  int    `json:"customer_id"`
	UserID      int    `json:"user_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type OrderUpdateRequest struct {
	CustomerID  int    `json:"customer_id"`
	UserID      int    `json:"user_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
}

type OrderModel struct {
}

func NewOrderModel() *OrderModel {
	return new(OrderModel)
}

func (o *OrderModel) GetAllOrders() ([]Order, error) {
	rows, err := database.Turso.Query("SELECT order_id, customer_id, user_id, name, description, order_date, status FROM order")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orders := []Order{}
	for rows.Next() {
		order := Order{}
		err := rows.Scan(&order.ID, &order.CustomerID, &order.UserID, &order.Name, &order.Description, &order.OrderDate, &order.Status)
		if err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}

func (o *OrderModel) CreateOrder(data *OrderCreateRequest) (*Order, error) {
	stmt, err := database.Turso.Prepare("INSERT INTO user (customer_id, user_id, name, description) VALUES (?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	exec, err := stmt.Exec(data.CustomerID, data.UserID, data.Name, data.Description)
	if err != nil {
		return nil, err
	}

	id, err := exec.LastInsertId()
	if err != nil {
		return nil, err
	}

	order := &Order{
		ID:          id,
		CustomerID:  data.CustomerID,
		UserID:      data.UserID,
		Name:        data.Name,
		Description: data.Description,
	}

	return order, nil
}

func (o *OrderModel) GetOrder(id int) (*Order, error) {
	order := new(Order)
	err := database.Turso.QueryRow(
		`SELECT order_id, customer_id, user_id, name, description, order_date, status
		FROM order
		WHERE order_id=?`,
		id,
	).Scan(
		&order.ID,
		&order.CustomerID,
		&order.UserID,
		&order.Name,
		&order.Description,
		&order.OrderDate,
		&order.Status,
	)

	if err != nil {
		return nil, err
	}

	return order, nil
}

func (o *OrderModel) UpdateOrder(id int, data *OrderUpdateRequest) (*Order, error) {
	stmt, err := database.Turso.Prepare("UPDATE order SET customer_id=?, user_id=?, name=?, description=?, status=? WHERE order_id=?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(data.CustomerID, data.UserID, data.Name, data.Description, data.Status, id)
	if err != nil {
		return nil, err
	}

	order, err := o.GetOrder(id)
	if err != nil {
		return nil, err
	}

	return order, nil
}
