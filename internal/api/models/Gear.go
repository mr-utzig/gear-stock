package models

import "database/sql"

type Gear struct {
	ID          int    `json:"id"`
	OrderID     int    `json:"order_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
}

type GearModel struct {
	db *sql.DB
}

func NewGearModel(db *sql.DB) *GearModel {
	return &GearModel{db: db}
}

func (o *GearModel) GetAllGears() *[]Gear {
	return &[]Gear{}
}

func (o *GearModel) CreateGear() *Gear {
	return &Gear{}
}

func (o *GearModel) GetGear(id int) *Gear {
	return &Gear{}
}

func (o *GearModel) UpdateGear(id int) *Gear {
	return &Gear{}
}

func (o *GearModel) DeleteGear(id int) *Gear {
	return &Gear{}
}
