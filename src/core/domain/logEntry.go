package domain

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type LogEntry struct {
	//Object metadata
	ID        primitive.ObjectID
	CreatedAt time.Time
	UpdatedAt time.Time

	//Identification (for reproduction)
	ProjectID      uuid.UUID
	Identification LogIdentification
	Data           LogData
}

type LogData struct {
	Timestamp        time.Time
	GroupingID       string
	Fingerprint      string
	Level            string
	Trace            Traceback
	NestedTrace      []Traceback
	Message          string
	AdditionalFields map[string]interface{}
}

type LogClientIdentification struct {
	UserID    *uuid.UUID
	IPAddress string
}

type LogDeploymentIdentification struct {
	ServerHostname string
	Environment    string
	Version        string
}

type LogIdentification struct {
	Client     LogClientIdentification
	Deployment LogDeploymentIdentification
}
