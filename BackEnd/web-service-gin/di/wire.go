//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	http "github.com/luke92/FriendsLessonsSystem/http"
	handler "github.com/luke92/FriendsLessonsSystem/handler"
	usecase "github.com/luke92/FriendsLessonsSystem/usecase"
	repository "github.com/luke92/FriendsLessonsSystem/usecase/repository"
)

func InitializeAPI() (*http.ServerHTTP, error) {
	wire.Build(repository.NewUserRepository, usecase.NewUserUseCase, handler.NewUserHandler, http.NewServerHTTP)

	return &http.ServerHTTP{}, nil
}
