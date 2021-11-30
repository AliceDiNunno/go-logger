package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (rH RoutesHandler) GetEnvironmentHandler(c *gin.Context) {
	user := rH.getAuthenticatedUser(c)
	if user == nil {
		return
	}

	project := rH.getProject(c)
	if project == nil {
		return
	}

	environments, err := rH.usecases.FetchProjectEnvironments(project)

	if err != nil {
		rH.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, environments)
}
