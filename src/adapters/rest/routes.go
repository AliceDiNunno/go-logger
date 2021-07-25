package rest

import (
	"github.com/gin-gonic/gin"
)

func SetRoutes(r *gin.Engine, routesHandler RoutesHandler) {
	r.Use(routesHandler.fetchingUserMiddleware())
	r.NoRoute(routesHandler.endpointNotFound)

	r.GET("", routesHandler.GetAllAppsHandler) //Get all apps
	r.POST("", routesHandler.CreateAppHandler) //Create an app

	appGroup := r.Group(":app_id", routesHandler.fetchingAppMiddleware())
	appGroup.GET("", routesHandler.GetAppHandler)       // Get an app
	appGroup.DELETE("", routesHandler.DeleteAppHandler) //Delete an app

	environmentGroup := appGroup.Group("/environment")
	environmentGroup.GET("", routesHandler.GetEnvironmentHandler) //Getting all declared environments for an app

	versionGroup := appGroup.Group("/version")
	versionGroup.GET("", routesHandler.GetVersionHandler) //Getting all declared version for an app

	serverGroup := appGroup.Group("/server")
	serverGroup.GET("", routesHandler.GetServerHandler) //Getting all declared servers for an app

	itemsGroup := appGroup.Group("/items")
	itemsGroup.GET("", routesHandler.SearchGroupingIdsHandler) //Search all grouping ids
	itemsGroup.POST("", routesHandler.PushLogsHandler)         //Push a log

	logsGroup := itemsGroup.Group("/:grouping_id", routesHandler.fetchingGroupMiddleware())
	logsGroup.GET("/", routesHandler.SearchLogsInGroupingHandler)   //Search all logs (corresponding to a grouping ID)
	logsGroup.GET("/:log_id", routesHandler.GetSpecificLogsHandler) //Getting a specific log id
}
