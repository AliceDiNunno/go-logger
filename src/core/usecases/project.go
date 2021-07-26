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

	album, err := i.ProjectRepo.FindById(user.UserID, id)

	if err != nil {
		return nil, err
	}

	return album, nil
}

func (i interactor) CreateProject(user *domain.User, createRequest *request.CreateProjectRequest) error {
	project := &domain.Project{
		Name:          createRequest.Name,
		RepositoryURL: createRequest.Url,
	}

	project.Initialize()

	projectSameName, err := i.ProjectRepo.FindByName(user.UserID, createRequest.Name)

	if err == nil && projectSameName != nil {
		//A workflow with the same name already exists for this user
		return domain.ErrProjectAlreadyExistingWithThisName
	}

	err = i.ProjectRepo.CreateProject(project)

	if err != nil {
		return err
	}

	return nil
}
