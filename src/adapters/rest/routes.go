package rest

import (
	"github.com/gin-gonic/gin"
)

func SetRoutes(r *gin.Engine, routesHandler RoutesHandler) {
	r.Use(routesHandler.fetchingUserMiddleware())
	r.NoRoute(routesHandler.endpointNotFound)

	r.GET("", routesHandler.GetAllAppsHandler) //Get all apps
	r.POST("", routesHandler.CreateAppHandler) //Create an app

	appGroup := r.Group(":id", routesHandler.fetchingAppMiddleware())
	appGroup.GET("", routesHandler.GetAppHandler)       // Get an app
	appGroup.DELETE("", routesHandler.DeleteAppHandler) //Delete an app

	environmentGroup := appGroup.Group("/environment")
	environmentGroup.GET("", nil) //Getting all declared environments for an app

	versionGroup := appGroup.Group("/version")
	versionGroup.GET("", nil) //Getting all declared version for an app

	serverGroup := appGroup.Group("/server")
	serverGroup.GET("", nil) //Getting all declared servers for an app

	itemsGroup := appGroup.Group("/items")
	itemsGroup.GET("", nil)    //Getting all grouping ids
	itemsGroup.GET(":id", nil) //Retrieve all logs (corresponding to a grouping ID)
	itemsGroup.POST("", nil)   //Push a log

}
