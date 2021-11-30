package postgres

import (
	"github.com/AliceDiNunno/go-logger/src/core/domain"
	"gorm.io/gorm"
)
import "github.com/google/uuid"

type Project struct {
	gorm.Model
	ID            uuid.UUID `gorm:"type:uuid;primary_key"`
	Name          string
	ProjectKey    uuid.UUID `gorm:"type:uuid;unique"`
	RepositoryURL string
	User          uuid.UUID
}

type projectRepo struct {
	db *gorm.DB
}

func (p projectRepo) FindByUser(user uuid.UUID) ([]*domain.Project, error) {
	var projects []*Project

	query := p.db.Where("\"user\" = ?", user).Find(&projects)

	if query.Error != nil {
		return nil, query.Error
	}

	return projectsToDomain(projects), nil
}

func (p projectRepo) FindByName(user uuid.UUID, name string) (*domain.Project, error) {
	var project *Project

	query := p.db.Where("\"user\" = ? AND name = ?", user, name).First(&project)

	if query.Error != nil {
		return nil, query.Error
	}

	return projectToDomain(project), nil
}

func (p projectRepo) CreateProject(project *domain.Project) error {
	projectToCreate := projectFromDomain(project)

	result := p.db.Create(projectToCreate)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (p projectRepo) FindById(user uuid.UUID, projectID uuid.UUID) (*domain.Project, error) {
	var project *Project

	query := p.db.Where("\"user\" = ? AND \"id\" = ?", user, projectID).First(&project)

	if query.Error != nil {
		return nil, query.Error
	}

	return projectToDomain(project), nil
}

func (p projectRepo) FindByIdAndKey(projectID uuid.UUID, projectKey uuid.UUID) (*domain.Project, error) {
	var project *Project

	query := p.db.Where("\"id\" = ? AND \"project_key\" = ?", projectID, projectKey).First(&project)

	if query.Error != nil {
		return nil, query.Error
	}

	return projectToDomain(project), nil
}

func (p projectRepo) DeleteProject(project *domain.Project) error {
	idToRemove := project.ID

	query := p.db.Where("id = ?", idToRemove).Delete(&Project{})

	return query.Error
}

func projectToDomain(project *Project) *domain.Project {
	return &domain.Project{
		ID:            project.ID,
		Name:          project.Name,
		ProjectKey:    project.ProjectKey,
		RepositoryURL: project.RepositoryURL,
		User:          project.User,
	}
}

func projectFromDomain(project *domain.Project) *Project {
	return &Project{
		ID:            project.ID,
		Name:          project.Name,
		ProjectKey:    project.ProjectKey,
		RepositoryURL: project.RepositoryURL,
		User:          project.User,
	}
}

func projectsToDomain(projects []*Project) []*domain.Project {
	projectList := []*domain.Project{}

	for _, project := range projects {
		projectList = append(projectList, projectToDomain(project))
	}

	return projectList
}

func NewProjectRepo(db *gorm.DB) projectRepo {
	return projectRepo{
		db: db,
	}
}
