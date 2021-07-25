package domain

import "github.com/google/uuid"

type App struct {
	ID     uuid.UUID
	Name   string
	AppKey uuid.UUID
}
