package rest

import (
	"github.com/AliceDiNunno/go-logger/src/core/domain/request"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func (rH RoutesHandler) fetchingGroupMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := rH.getAuthenticatedUser(c)
		if user == nil {
			return
		}

		project := rH.getProject(c)
		if project == nil {
			return
		}

		id := c.Param("grouping_id")

		if id == "" {
			rH.handleError(c, ErrFormValidation)
			return
		}

		c.Set("grouping", id)
	}
}

func (rH RoutesHandler) getGrouping(c *gin.Context) string {
	grouping, exists := c.Get("grouping")

	if !exists {
		return ""
	}

	foundGrouping := grouping.(string)

	return foundGrouping
}

func (rH RoutesHandler) PushLogsHandler(c *gin.Context) {
	var creationRequest request.ItemCreationRequest

	err := c.ShouldBind(&creationRequest)

	if err != nil {
		rH.handleError(c, ErrFormValidation)
		return
	}

	id, err := uuid.Parse(c.Param("project_id"))

	if err != nil {
		rH.handleError(c, ErrFormValidation)
		return
	}

	err = rH.usecases.PushNewLogEntry(id, &creationRequest)

	if err != nil {
		rH.handleError(c, err)
		return
	}

	c.Status(http.StatusCreated)
}

func (rH RoutesHandler) SearchLogsInGroupingHandler(c *gin.Context) {
	user := rH.getAuthenticatedUser(c)
	if user == nil {
		return
	}

	project := rH.getProject(c)
	if project == nil {
		return
	}

	groupingId := c.Param("grouping_id")

	logs, err := rH.usecases.FetchGroupingIdContent(project, groupingId)

	if err != nil {
		rH.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, logs)
}

func (rH RoutesHandler) GetLogsOccurencesHandler(c *gin.Context) {
	user := rH.getAuthenticatedUser(c)
	if user == nil {
		return
	}

	project := rH.getProject(c)
	if project == nil {
		return
	}

	groupingId := rH.getGrouping(c)
	if groupingId == "" {
		return
	}

	logs, err := rH.usecases.FetchGroupingIdOccurrences(project, groupingId)

	if err != nil {
		rH.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, logs)
}

func (rH RoutesHandler) GetSpecificLogsHandler(c *gin.Context) {
	user := rH.getAuthenticatedUser(c)
	if user == nil {
		return
	}

	project := rH.getProject(c)
	if project == nil {
		return
	}

	groupingId := rH.getGrouping(c)
	if groupingId == "" {
		return
	}

	logId := c.Param("log_id")
	if logId == "" {
		rH.handleError(c, ErrFormValidation)
	}

	logs, err := rH.usecases.FetchGroupOccurrence(project, groupingId, logId)

	if err != nil {
		rH.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, logs)
}
