package rest

import (
	"github.com/AliceDiNunno/go-logger/src/core/domain/request"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func (rH RoutesHandler) fetchingGroupMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func (rH RoutesHandler) SearchGroupingIdsHandler(c *gin.Context) {

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

}

func (rH RoutesHandler) GetSpecificLogsHandler(c *gin.Context) {

}
