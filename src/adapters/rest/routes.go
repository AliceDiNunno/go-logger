package rest

func SetRoutes(server GinServer, routesHandler RoutesHandler) {
	r := server.Router

	r.NoRoute(routesHandler.endpointNotFound)

	main := r.Group(server.Config.Prefix, routesHandler.fetchingUserMiddleware())

	main.GET("", routesHandler.GetUserProjectsHandler) //Get all projects
	main.POST("", routesHandler.CreateProjectHandler)  //Create a project

	projectGroup := main.Group(":project_id", routesHandler.fetchingProjectMiddleware())
	projectGroup.DELETE("", routesHandler.DeleteProjectHandler) //Delete a project

	environmentGroup := projectGroup.Group("/environment")
	environmentGroup.GET("", routesHandler.GetEnvironmentHandler) //Getting all declared environments for a project

	versionGroup := projectGroup.Group("/version")
	versionGroup.GET("", routesHandler.GetVersionHandler) //Getting all declared version for a project

	serverGroup := projectGroup.Group("/server")
	serverGroup.GET("", routesHandler.GetServerHandler) //Getting all declared servers for a project

	itemsGroup := projectGroup.Group("/items")
	itemsGroup.GET("", routesHandler.GetProjectHandler)         //Search all grouping ids
	r.POST("/:project_id/items", routesHandler.PushLogsHandler) //Push a log

	logsGroup := itemsGroup.Group("/:grouping_id", routesHandler.fetchingGroupMiddleware())
	logsGroup.GET("/occurrences", routesHandler.GetLogsOccurencesHandler)       //Getting a specific log id
	logsGroup.GET("/", routesHandler.SearchLogsInGroupingHandler)               //Search all logs (corresponding to a grouping ID)
	logsGroup.GET("/occurrences/:log_id", routesHandler.GetSpecificLogsHandler) //Getting a specific log id
}
