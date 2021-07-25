package domain

import "github.com/google/uuid"

type Project struct {
	ID        uuid.UUID
	Name      string
	ProjectId uuid.UUID
	Repo      string
}
