package usecase

import (
	domain "github.com/luke92/FriendsLessonsSystem/domain"
	services "github.com/luke92/FriendsLessonsSystem/usecase/interface"
)

type userUseCase struct {
	userRepo domain.UserRepository
}

func NewUserUseCase(repo domain.UserRepository) services.UserUseCase {
	return &userUseCase{
		userRepo: repo,
	}
}

func (c *userUseCase) GetAll() ([]domain.User, error) {
	users, err := c.userRepo.GetAll()
	return users, err
}

func (c *userUseCase) GetByID(id string) (domain.User, error) {
	user, err := c.userRepo.GetByID(id)
	return user, err
}
