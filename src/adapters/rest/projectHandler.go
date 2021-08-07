package rest

import (
	"github.com/AliceDiNunno/go-logger/src/core/domain"
	"github.com/AliceDiNunno/go-logger/src/core/domain/request"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func (rH RoutesHandler) fetchingProjectMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		println("Check auth")
		user := rH.getAuthenticatedUser(c)
		if user == nil {
			return
		}

		id, err := uuid.Parse(c.Param("project_id"))

		if err != nil {
			rH.handleError(c, ErrFormValidation)
			return
		}

		project, err := rH.usecases.FetchProject(user, id)

		if err != nil {
			rH.handleError(c, domain.ErrProjectNotFound)
			return
		}

		if project.User != user.UserID {
			rH.handleError(c, domain.ErrProjectNotFound)
			return
		}

		c.Set("project", project)
	}
}

func (rH RoutesHandler) getProject(c *gin.Context) *domain.Project {
	project, exists := c.Get("project")

	if !exists {
		return nil
	}

	foundProject := project.(*domain.Project)

	return foundProject
}

func (rH RoutesHandler) GetUserProjectsHandler(c *gin.Context) {
	user := rH.getAuthenticatedUser(c)
	if user == nil {
		return
	}

	albums, err := rH.usecases.GetUserProjects(user)

	if err != nil {
		rH.handleError(c, err)
	}

	c.JSON(http.StatusOK, albums)
}

func (rH RoutesHandler) CreateProjectHandler(c *gin.Context) {
	user := rH.getAuthenticatedUser(c)

	var request request.CreateProjectRequest
	err := c.ShouldBind(&request)
	if err != nil {
		rH.handleError(c, ErrFormValidation)
		return
	}

	err = rH.usecases.CreateProject(user, &request)

	if err != nil {
		rH.handleError(c, err)
		return
	}

	c.Status(http.StatusCreated)
}

func (rH RoutesHandler) GetProjectHandler(c *gin.Context) {
	user := rH.getAuthenticatedUser(c)
	if user == nil {
		return
	}

	project := rH.getProject(c)
	if project == nil {
		return
	}

	result, err := rH.usecases.GetProjectsContent(user, project)

	if err != nil {
		rH.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (rH RoutesHandler) DeleteProjectHandler(c *gin.Context) {
	user := rH.getAuthenticatedUser(c)
	if user == nil {
		return
	}

	project := rH.getProject(c)
	if project == nil {
		return
	}

	err := rH.usecases.DeleteProject(user, project)

	if err != nil {
		rH.handleError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}
