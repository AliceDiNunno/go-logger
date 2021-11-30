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
	GetProjectsContent(user *domain.User, project *domain.Project) ([]string, error)
	GetUserProjects(user *domain.User) ([]*domain.Project, error)

	PushNewLogEntry(id uuid.UUID, request *request.ItemCreationRequest) error

	FetchGroupingIdContent(project *domain.Project, groupingId string) (*domain.LogEntry, error)
	FetchGroupingIdOccurrences(project *domain.Project, groupingId string) ([]string, error)
	FetchGroupOccurrence(project *domain.Project, groupingId string, occurrence string) (*domain.LogEntry, error)
	FetchProjectVersions(project *domain.Project) ([]string, error)
	FetchProjectEnvironments(project *domain.Project) ([]string, error)
	FetchProjectServers(project *domain.Project) ([]string, error)
}
