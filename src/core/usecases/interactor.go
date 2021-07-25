package usecases

import "github.com/AliceDiNunno/go-logger/src/core/domain"

type Logger interface {
	Error(args ...interface{})
	Fatal(args ...interface{})
	Info(args ...interface{})
	Debug(args ...interface{})
}

type AppRepo interface {
}

type LogCollection interface {
	AddLog(log *domain.LogEntry) error
}

type interactor struct {
	AppRepo       AppRepo
	LogCollection LogCollection
}

func NewInteractor(aR AppRepo, lC LogCollection) interactor {
	return interactor{
		AppRepo:       aR,
		LogCollection: lC,
	}
}
