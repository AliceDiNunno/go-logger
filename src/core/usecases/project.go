package usecases

import (
	"github.com/AliceDiNunno/go-logger/src/core/domain"
	"github.com/AliceDiNunno/go-logger/src/core/domain/request"
	"github.com/google/uuid"
)

func (i interactor) FetchProject(user *domain.User, id uuid.UUID) (*domain.Project, error) {
	if user == nil {
		return nil, domain.ErrFailedToGetUser
	}

	album, err := i.projectRepo.FindById(user.UserID, id)

	if err != nil {
		return nil, err
	}

	return album, nil
}

func (i interactor) CreateProject(user *domain.User, createRequest *request.CreateProjectRequest) error {
	project := &domain.Project{
		Name:          createRequest.Name,
		RepositoryURL: createRequest.Url,
		User:          user.UserID,
	}

	project.Initialize()

	projectSameName, err := i.projectRepo.FindByName(user.UserID, createRequest.Name)

	if err == nil && projectSameName != nil {
		//A workflow with the same name already exists for this user
		return domain.ErrProjectAlreadyExistingWithThisName
	}

	err = i.projectRepo.CreateProject(project)

	if err != nil {
		return err
	}

	return nil
}

func (i interactor) DeleteProject(user *domain.User, project *domain.Project) error {
	err := i.projectRepo.DeleteProject(project)

	if err != nil {
		return domain.ErrUnableToDeleteObject
	}

	return nil
}

func (i interactor) GetProjectsContent(user *domain.User, project *domain.Project) ([]string, error) {
	if user == nil {
		return nil, domain.ErrFailedToGetUser
	}
	return i.logCollection.ProjectGroupingIds(project)
}

func (i interactor) GetUserProjects(user *domain.User) ([]*domain.Project, error) {
	workflows, err := i.projectRepo.FindByUser(user.UserID)

	if err != nil {
		return nil, err
	}

	return workflows, nil
}
