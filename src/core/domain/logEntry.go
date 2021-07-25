package domain

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type LogEntry struct {
	ID          primitive.ObjectID `bson:"_id"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
	Trace       Traceback          `bson:"trace"`
	Environment string             `bson:"environment"`
	Project     uuid.UUID          `bson:"project"`
}
