package entities

import "github.com/google/uuid"

type Restaurant struct {
	ID       uuid.UUID
	Name     string
	Location string
}
