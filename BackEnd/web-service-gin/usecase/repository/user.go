package repository

import (
	domain "github.com/luke92/FriendsLessonsSystem/domain"
)

type userDatabase struct {
	users []domain.User
}

func NewUserRepository() domain.UserRepository {
	var users = []domain.User{
		{Id: "1", Username: "luke92", Password: "mypassword", Email: "lucasjv92@gmail.com", Name: "Lucas", Lastname: "Vargas"},
		{Id: "2", Username: "henry", Password: "henryPass", Email: "henry@modak.live", Name: "Henry", Lastname: "Canastero"},
	}
	return &userDatabase{users}
}

func (c *userDatabase) GetAll() ([]domain.User, error) {
	return c.users, nil
}

func (c *userDatabase) GetByID(id string) (domain.User, error) {
	var user domain.User
	// Loop through the list of users, looking for
	// an user whose ID value matches the parameter.
	for i := range c.users {
		if c.users[i].Id == id {
			user = c.users[i]
			break
		}
	}

	return user, nil
}
