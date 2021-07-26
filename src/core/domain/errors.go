package domain

import "errors"

var (
	ErrFailedToGetUser                    = errors.New("failed to fetch user")
	ErrProjectNotFound                    = errors.New("project not found")
	ErrProjectAlreadyExistingWithThisName = errors.New("a project already exists with this name")
)
