package usecases

import "github.com/AliceDiNunno/go-logger/src/core/domain"

func (i interactor) FetchProjectVersions(project *domain.Project) ([]string, error) {
	return i.logCollection.ProjectVersions(project)
}
