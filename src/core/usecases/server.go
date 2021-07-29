package usecases

import "github.com/AliceDiNunno/go-logger/src/core/domain"

func (i interactor) FetchProjectServers(project *domain.Project) ([]string, error) {
	return i.logCollection.ProjectServers(project)

}
