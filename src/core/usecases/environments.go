package usecases

import "github.com/AliceDiNunno/go-logger/src/core/domain"

func (i interactor) FetchProjectEnvironments(project *domain.Project) ([]string, error) {
	environments, err := i.logCollection.ProjectEnvironments(project)

	if err != nil {
		return nil, domain.ErrUnknownDBError
	}

	return environments, nil
}
