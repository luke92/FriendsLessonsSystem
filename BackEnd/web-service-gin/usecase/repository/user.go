package repository

import (
	domain "github.com/luke92/FriendsLessonsSystem/domain"
)

type userDatabase struct {
	users []domain.User
}

func NewUserRepository() domain.UserRepository {
	var users = []domain.User{
		{Id: "1", Username: "joe", Password: "joepass", Email: "joe@gmail.com", Name: "Joe", Lastname: "Reid"},
		{Id: "2", Username: "mark", Password: "markPass", Email: "mark@gmail.com", Name: "Mark", Lastname: "Kennedy"},
		{Id: "3", Username: "jody", Password: "jodyPass", Email: "jody@gmail.com", Name: "Jody", Lastname: "Wagner"},
		{Id: "4", Username: "rachel", Password: "rachelPass", Email: "rachel@gmail.com", Name: "Rachel", Lastname: "Miles"},
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
