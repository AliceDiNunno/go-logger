package domain

import "github.com/google/uuid"

type Occurrence struct {
	ID          uuid.UUID
	Server      *Server
	Error       *error
	Origin      string
	User        *uuid.UUID
	Environment *Environment
	Traceback   Traceback
}
