package handler

import (
	"github.com/gin-gonic/gin"

	service "producer/service"
)

type Handlers struct {
	service service.Service
}

func New(service service.Service) *Handlers {
	return &Handlers{
		service: service,
	}
}

func (app *Handlers) RegisterRoutes(router *gin.Engine) *Handlers {
	
	router.POST("sender", app.Sender)
	return app
}
