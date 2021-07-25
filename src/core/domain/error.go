package domain

import "github.com/google/uuid"

type Error struct {
	ID          uuid.UUID
	Number      int
	Application *Project
	Name        string
	Resolved    bool
	Muted       bool
}
