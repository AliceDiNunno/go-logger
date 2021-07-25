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
	ProjectID      uuid.UUID          `bson:"project"`
	ServerHostname string             `bson:"server_hostname"`
	Environment    string             `bson:"environment"`
	Level          string             `bson:"level"`
	Version        string             `bson:"version"`
	Trace          domain.Traceback   `bson:"trace"`
	NestedTrace    []domain.Traceback `bson:"nested_trace"`
	UserID         *uuid.UUID         `bson:"user_id"`
	IPAddress      string             `bson:"ip_address"`

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

func (c logCollection) AddLog(entry *domain.LogEntry) error {
	_, err := c.collection.InsertOne(context.Background(), entry)
	return err
}

func NewLogCollectionRepo(db *mongo.Client) logCollection {
	collection := db.Database("logger").Collection("logs")

	return logCollection{
		db:         db,
		collection: collection,
	}
}
