package usecases

import "github.com/AliceDiNunno/go-logger/src/core/domain"

func (i interactor) FetchProjectEnvironments(project *domain.Project) ([]string, error) {
	return i.logCollection.ProjectEnvironments(project)
}
