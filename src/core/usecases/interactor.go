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
	FindByIdAndKey(project uuid.UUID, projectKey uuid.UUID) (*domain.Project, error)
	FindByName(user uuid.UUID, name string) (*domain.Project, error)
	FindByUser(user uuid.UUID) ([]*domain.Project, error)
	DeleteProject(project *domain.Project) error
}

type LogCollection interface {
	AddLog(log *domain.LogEntry) error

	ProjectVersions(project *domain.Project) ([]string, error)
	ProjectEnvironments(project *domain.Project) ([]string, error)
	ProjectServers(project *domain.Project) ([]string, error)
	ProjectGroupingIds(project *domain.Project) ([]string, error)

	FindLastEntryForGroup(project *domain.Project, groupingId string) (*domain.LogEntry, error)
	FindGroupOccurrences(project *domain.Project, groupingId string) ([]string, error)
	FindGroupOccurrence(project *domain.Project, groupingId string, occurenceId string) (*domain.LogEntry, error)
}

type interactor struct {
	projectRepo   ProjectRepo
	logCollection LogCollection
}

func NewInteractor(pR ProjectRepo, lC LogCollection) interactor {
	return interactor{
		projectRepo:   pR,
		logCollection: lC,
	}
}
