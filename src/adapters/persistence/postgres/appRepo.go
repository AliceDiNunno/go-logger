package postgres

import "gorm.io/gorm"
import "github.com/google/uuid"

type App struct {
	gorm.Model
	ID            uuid.UUID `gorm:"type:uuid;primary_key"`
	Name          string
	AppKey        uuid.UUID `gorm:"type:uuid;unique"`
	RepositoryURL string
}

type appRepo struct {
	db *gorm.DB
}

func NewAppRepo(db *gorm.DB) appRepo {
	return appRepo{
		db: db,
	}
}
