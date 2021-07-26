package usecases

import (
	"github.com/AliceDiNunno/go-logger/src/core/domain"
	"github.com/AliceDiNunno/go-logger/src/core/domain/request"
	"github.com/google/uuid"
)

type Usecases interface {
	FetchProject(user *domain.User, id uuid.UUID) (*domain.Project, error)
	CreateProject(user *domain.User, createRequest *request.CreateProjectRequest) error
	DeleteProject(user *domain.User, project *domain.Project) error
	GetAlbumsContent(user *domain.User, project *domain.Project)
	GetUserProjects(user *domain.User) ([]*domain.Project, error)
}
