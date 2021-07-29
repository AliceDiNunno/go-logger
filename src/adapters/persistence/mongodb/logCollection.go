package mongodb

import (
	"context"
	"github.com/AliceDiNunno/go-logger/src/core/domain"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type logEntry struct {
	//Object metadata
	ID        primitive.ObjectID `bson:"_id"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`

	//Entry metadata
	Timestamp   time.Time `bson:"timestamp"`
	GroupingID  string    `bson:"grouping_id"`
	Fingerprint string    `bson:"fingerprint"`

	//Identification (for reproduction)
	Platform    string             `bson:"platform"`
	Source      string             `bson:"source"`
	ProjectID   uuid.UUID          `bson:"project"`
	Hostname    string             `bson:"hostname"`
	Environment string             `bson:"environment"`
	Level       string             `bson:"level"`
	Version     string             `bson:"version"`
	Trace       domain.Traceback   `bson:"trace"`
	NestedTrace []domain.Traceback `bson:"nested_trace"`
	UserID      *uuid.UUID         `bson:"user_id"`
	IPAddress   string             `bson:"ip_address"`
	StatusCode  int                `bson:"status_code"`

	//Entry Data
	Message          string                 `bson:"message"`
	AdditionalFields map[string]interface{} `bson:"additional"`
}

type SearchLogsFilter struct {
	ProjectID      *uuid.UUID
	UserID         *uuid.UUID
	Fingerprint    string
	GroupingID     string
	ServerHostname string
	Environment    string
	Level          string
	IPAddress      string
}

type logCollection struct {
	db         *mongo.Client
	collection *mongo.Collection
}

func logEntryToDomain(entry *logEntry) *domain.LogEntry {
	return &domain.LogEntry{
		ID:        entry.ID,
		CreatedAt: entry.CreatedAt,
		UpdatedAt: entry.UpdatedAt,
		ProjectID: entry.ProjectID,
		Identification: domain.LogIdentification{
			Client: domain.LogClientIdentification{
				UserID:    entry.UserID,
				IPAddress: entry.IPAddress,
			},
			Deployment: domain.LogDeploymentIdentification{
				Platform:    entry.Platform,
				Source:      entry.Source,
				Hostname:    entry.Hostname,
				Environment: entry.Environment,
				Version:     entry.Version,
			},
		},
		Data: domain.LogData{
			Timestamp:        entry.Timestamp,
			GroupingID:       entry.GroupingID,
			Fingerprint:      entry.Fingerprint,
			Level:            entry.Level,
			Trace:            entry.Trace,
			NestedTrace:      entry.NestedTrace,
			Message:          entry.Message,
			StatusCode:       entry.StatusCode,
			AdditionalFields: entry.AdditionalFields,
		},
	}
}

func logEntryFromDomain(entry *domain.LogEntry) *logEntry {
	return &logEntry{
		ID:               entry.ID,
		CreatedAt:        entry.CreatedAt,
		UpdatedAt:        entry.UpdatedAt,
		ProjectID:        entry.ProjectID,
		Timestamp:        entry.Data.Timestamp,
		GroupingID:       entry.Data.GroupingID,
		Fingerprint:      entry.Data.Fingerprint,
		Level:            entry.Data.Level,
		Trace:            entry.Data.Trace,
		AdditionalFields: entry.Data.AdditionalFields,
		NestedTrace:      entry.Data.NestedTrace,
		StatusCode:       entry.Data.StatusCode,
		Message:          entry.Data.Message,
		Platform:         entry.Identification.Deployment.Platform,
		Source:           entry.Identification.Deployment.Source,
		Hostname:         entry.Identification.Deployment.Hostname,
		Environment:      entry.Identification.Deployment.Environment,
		Version:          entry.Identification.Deployment.Version,
		UserID:           entry.Identification.Client.UserID,
		IPAddress:        entry.Identification.Client.IPAddress,
	}
}

func (c logCollection) AddLog(entry *domain.LogEntry) error {
	entryFromDomain := logEntryFromDomain(entry)

	_, err := c.collection.InsertOne(context.Background(), entryFromDomain)

	return err
}

func NewLogCollectionRepo(db *mongo.Client) logCollection {
	collection := db.Database("logger").Collection("logs")

	return logCollection{
		db:         db,
		collection: collection,
	}
}
