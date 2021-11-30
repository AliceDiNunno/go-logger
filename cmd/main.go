package main

import (
	"fmt"
	"github.com/AliceDiNunno/go-logger/src/adapters/persistence/mongodb"
	"github.com/AliceDiNunno/go-logger/src/adapters/persistence/postgres"
	"github.com/AliceDiNunno/go-logger/src/adapters/rest"
	"github.com/AliceDiNunno/go-logger/src/config"
	"github.com/AliceDiNunno/go-logger/src/core/usecases"
	"gorm.io/gorm"
	"log"
)

func main() {
	config.LoadEnv()

	ginConfiguration := config.LoadGinConfiguration()
	dbConfig := config.LoadGormConfiguration()
	mongoConfig := config.LoadMongodbConfiguration()

	mongo := mongodb.StartMongodbDatabase(mongoConfig)
	var logCollection usecases.LogCollection

	logCollection = mongodb.NewLogCollectionRepo(mongo)

	var projectRepo usecases.ProjectRepo
	var db *gorm.DB
	if dbConfig.Engine == "POSTGRES" {
		db = postgres.StartGormDatabase(dbConfig)
		projectRepo = postgres.NewProjectRepo(db)

		err := db.AutoMigrate(&postgres.Project{})
		if err != nil {
			log.Fatalln(err)
		}
	} else {
		log.Fatalln(fmt.Sprintf("Database engine \"%s\" not supported", dbConfig.Engine))
	}

	usecasesHandler := usecases.NewInteractor(projectRepo, logCollection)

	restServer := rest.NewServer(ginConfiguration)
	routesHandler := rest.NewRouter(usecasesHandler)

	rest.SetRoutes(restServer.Router, routesHandler)

	restServer.Start()
}
