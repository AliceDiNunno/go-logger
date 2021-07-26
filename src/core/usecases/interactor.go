package usecases

import (
	"github.com/AliceDiNunno/go-logger/src/core/domain"
	"github.com/google/uuid"
)

type Logger interface {
	Error(args ...interface{})
	Fatal(args ...interface{})
	Info(args ...interface{})
	Debug(args ...interface{})
}

type ProjectRepo interface {
	CreateProject(project *domain.Project) error
	FindById(user uuid.UUID, project uuid.UUID) (*domain.Project, error)
	FindByName(user uuid.UUID, name string) (*domain.Project, error)
	FindByUser(user uuid.UUID) ([]*domain.Project, error)
	DeleteProject(project *domain.Project) error
}

type LogCollection interface {
	AddLog(log *domain.LogEntry) error
}

type interactor struct {
	ProjectRepo   ProjectRepo
	LogCollection LogCollection
}

func (i interactor) DeleteProject(user *domain.User, project *domain.Project) error {
	panic("implement me")
}

func (i interactor) GetAlbumsContent(user *domain.User, project *domain.Project) {
	panic("implement me")
}

func (i interactor) GetUserProjects(user *domain.User) ([]*domain.Project, error) {
	panic("implement me")
}

func NewInteractor(pR ProjectRepo, lC LogCollection) interactor {
	return interactor{
		ProjectRepo:   pR,
		LogCollection: lC,
	}
}
