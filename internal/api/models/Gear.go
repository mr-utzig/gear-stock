package models

import "github.com/mr-utzig/gear-stock/internal/database"

type Gear struct {
	ID          int64  `json:"id"`
	OrderID     int    `json:"order_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
}

type GearCreateRequest struct {
	OrderID     int    `json:"order_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
}

type GearUpdateRequest struct {
	OrderID     int    `json:"order_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
}

type GearModel struct {
}

func NewGearModel() *GearModel {
	return new(GearModel)
}

func (g *GearModel) GetAllGears() ([]Gear, error) {
	rows, err := database.Turso.Query("SELECT gear_id, order_id, name, description, status FROM gear")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	gears := []Gear{}
	for rows.Next() {
		gear := Gear{}
		err := rows.Scan(&gear.ID, &gear.OrderID, &gear.Name, &gear.Description, &gear.Status)
		if err != nil {
			return nil, err
		}

		gears = append(gears, gear)
	}

	return gears, nil
}

func (g *GearModel) CreateGear(data *GearCreateRequest) (*Gear, error) {
	stmt, err := database.Turso.Prepare("INSERT INTO gear (order_id, name, description, status) VALUES (?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	exec, err := stmt.Exec(data.OrderID, data.Name, data.Description, data.Status)
	if err != nil {
		return nil, err
	}

	id, err := exec.LastInsertId()
	if err != nil {
		return nil, err
	}

	gear := &Gear{
		ID:          id,
		OrderID:     data.OrderID,
		Name:        data.Name,
		Description: data.Description,
		Status:      data.Status,
	}

	return gear, nil
}

func (g *GearModel) GetGear(id int) (*Gear, error) {
	gear := new(Gear)
	err := database.Turso.QueryRow(
		`SELECT gear_id, order_id, name, description, status
		FROM gear
		WHERE gear_id=?`,
		id,
	).Scan(
		&gear.ID,
		&gear.OrderID,
		&gear.Name,
		&gear.Description,
		&gear.Status,
	)

	if err != nil {
		return nil, err
	}

	return gear, nil
}

func (g *GearModel) UpdateGear(id int, data *GearUpdateRequest) (*Gear, error) {
	stmt, err := database.Turso.Prepare("UPDATE gear SET order_id=?, name=?, description=?, status=? WHERE gear_id=?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(data.OrderID, data.Name, data.Description, data.Status, id)
	if err != nil {
		return nil, err
	}

	gear, err := g.GetGear(id)
	if err != nil {
		return nil, err
	}

	return gear, nil
}

func (g *GearModel) DeleteGear(id int) (*int, error) {
	stmt, err := database.Turso.Prepare("DELETE FROM gear WHERE gear_id=?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return nil, err
	}

	return &id, nil
}
