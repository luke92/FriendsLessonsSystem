package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	services "github.com/luke92/FriendsLessonsSystem/usecase/interface"
)

type UserHandler struct {
	userUseCase services.UserUseCase
}

func NewUserHandler(usecase services.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: usecase,
	}
}

func (cr *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := cr.userUseCase.GetAll()

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, users)
	}
}

func (cr *UserHandler) GetAllFriendships(c *gin.Context) {
	friendships, err := cr.userUseCase.GetAllFriendships()

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, friendships)
	}
}

func (cr *UserHandler) GetUserByID(c *gin.Context) {
	id := c.Param("id")

	user, err := cr.userUseCase.GetByID(id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, user)
	}
}

func (cr *UserHandler) GetFriendsByUserID(c *gin.Context) {
	id := c.Param("id")

	friends, err := cr.userUseCase.GetFriendsByID(id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, friends)
	}
}

func (cr *UserHandler) GetLessonsByUserID(c *gin.Context) {
	id := c.Param("id")

	lessons, err := cr.userUseCase.GetLessonsByID(id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, lessons)
	}
}
