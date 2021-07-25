package domain

import "github.com/google/uuid"

type Environment struct {
	Id   uuid.UUID
	Name string
}
