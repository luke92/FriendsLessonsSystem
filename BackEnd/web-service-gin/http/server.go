package http

import (
	"github.com/gin-gonic/gin"
	handler "github.com/luke92/FriendsLessonsSystem/http/handler"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(userHandler *handler.UserHandler) *ServerHTTP {
	engine := gin.New()

	api := engine.Group("/api")
	api.GET("users", userHandler.GetAll)
	api.GET("users/:id", userHandler.GetByID)

	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":8080")
}
