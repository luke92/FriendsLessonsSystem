package services

import (
	domain "github.com/luke92/FriendsLessonsSystem/domain"
)

type UserUseCase interface {
	GetAll() ([]domain.User, error)
	GetByID(id string) (domain.User, error)
	GetAllFriendships() ([]string, error)
	GetFriendsByID(id string) ([]string, error)
	GetLessonsByID(id string) ([]domain.Lesson, error)
}
