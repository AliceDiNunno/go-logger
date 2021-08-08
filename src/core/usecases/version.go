package usecases

import "github.com/AliceDiNunno/go-logger/src/core/domain"

func (i interactor) FetchProjectVersions(project *domain.Project) ([]string, error) {
	versions, err := i.logCollection.ProjectVersions(project)

	if err != nil {
		return nil, domain.ErrUnknownDBError
	}

	return versions, nil
}
