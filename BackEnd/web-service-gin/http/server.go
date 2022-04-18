package http

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	handler "github.com/luke92/FriendsLessonsSystem/http/handler"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(userHandler *handler.UserHandler) *ServerHTTP {
	engine := gin.New()
	engine.Use(cors.Default())
	api := engine.Group("/api")
	api.GET("users", userHandler.GetAllUsers)
	api.GET("friendships", userHandler.GetAllFriendships)
	api.GET("users/:id", userHandler.GetUserByID)
	api.GET("users/:id/friends", userHandler.GetFriendsByUserID)
	api.GET("users/:id/lessons", userHandler.GetLessonsByUserID)

	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start(port string) {
	sh.engine.Run(":" + port)
}
