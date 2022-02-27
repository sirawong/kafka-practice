package handler

import (
	"events"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctrl *Handlers) Sender(c *gin.Context) {
	inputRequest := &events.Messages{}
	err := c.ShouldBind(inputRequest)
	if err != nil {
		HandleError(c, err)
	}

	msg, err := ctrl.service.Sender(inputRequest)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, msg)
}
