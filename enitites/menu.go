package entities

import "github.com/google/uuid"

type MenuItem struct {
	ID       uuid.UUID
	Name     string
	Price    float64
	Category string
}
