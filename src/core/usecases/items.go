package usecases

import (
	"github.com/AliceDiNunno/go-logger/src/core/domain"
	"github.com/AliceDiNunno/go-logger/src/core/domain/request"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func (i interactor) PushNewLogEntry(id uuid.UUID, request *request.ItemCreationRequest) error {
	project, error := i.projectRepo.FindByIdAndKey(id, request.ProjectKey)

	if error != nil || project == nil {
		return domain.ErrProjectNotFound
	}

	logEntry := &domain.LogEntry{
		ID:             primitive.NewObjectID(),
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		ProjectID:      id,
		Identification: request.Identification,
		Data:           request.Data,
	}

	if logEntry.Data.GroupingID == "" {
		logEntry.Data.GroupingID = logEntry.Data.Fingerprint
	}

	return i.logCollection.AddLog(logEntry)
}

func (i interactor) FetchGroupingIdContent(project *domain.Project, groupingId string) (*domain.LogEntry, error) {
	return i.logCollection.FindLastEntryForGroup(project, groupingId)
}

func (i interactor) FetchGroupingIdOccurrences(project *domain.Project, groupingId string) ([]string, error) {
	return i.logCollection.FindGroupOccurrences(project, groupingId)
}

func (i interactor) FetchGroupOccurrence(project *domain.Project, groupingId string, occurrence string) (*domain.LogEntry, error) {
	return i.logCollection.FindGroupOccurrence(project, groupingId, occurrence)
}
