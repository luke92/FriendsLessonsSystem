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

func (c *userUseCase) GetAllFriendships() ([]string, error) {
	var friendships []string
	users, err := c.userRepo.GetAll()

	for i := range users {
		var friends = users[i].Friends
		var name = users[i].Name
		if len(friends) > 0 {
			if !contains(friendships, name) {
				friendships = append(friendships, name)
				for j := range friends {
					friendships = append(friendships, friends[j].Name)
				}
			}
		}
	}

	return friendships, err
}

func (c *userUseCase) GetFriendsByID(id string) ([]string, error) {
	var friendships []string
	user, err := c.userRepo.GetByID(id)

	for i := range user.Friends {
		friendships = append(friendships, user.Friends[i].Name)
	}

	return friendships, err
}

func (c *userUseCase) GetLessonsByID(id string) ([]domain.Lesson, error) {
	var lessons []domain.Lesson
	user, err := c.userRepo.GetByID(id)

	lessons = user.Lessons

	return lessons, err
}

func contains(elems []string, v string) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}
