package domain

import "github.com/google/uuid"

type Project struct {
	ID            uuid.UUID
	Name          string
	ProjectKey    uuid.UUID
	RepositoryURL string
	User          uuid.UUID
}

func (p *Project) Initialize() {
	p.ID = uuid.New()
	p.ProjectKey = uuid.New()
}
