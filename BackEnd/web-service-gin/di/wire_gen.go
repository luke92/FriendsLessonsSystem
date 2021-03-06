// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/luke92/FriendsLessonsSystem/http"
	"github.com/luke92/FriendsLessonsSystem/http/handler"
	"github.com/luke92/FriendsLessonsSystem/usecase"
	"github.com/luke92/FriendsLessonsSystem/usecase/repository"
)

// Injectors from wire.go:

func InitializeAPI() (*http.ServerHTTP, error) {
	userRepository := repository.NewUserRepository()
	userUseCase := usecase.NewUserUseCase(userRepository)
	userHandler := handler.NewUserHandler(userUseCase)
	serverHTTP := http.NewServerHTTP(userHandler)
	return serverHTTP, nil
}
