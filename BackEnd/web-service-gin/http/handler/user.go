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

// FindAll godoc
// @summary Get all users
// @description Get all users
// @tags users
// @security ApiKeyAuth
// @id GetAll
// @produce json
// @Router /api/users [get]
// @response 200 {object} []Response "OK"
func (cr *UserHandler) GetAll(c *gin.Context) {
	users, err := cr.userUseCase.GetAll()

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, users)
	}
}

func (cr *UserHandler) GetByID(c *gin.Context) {
	id := c.Param("id")

	user, err := cr.userUseCase.GetByID(id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, user)
	}
}