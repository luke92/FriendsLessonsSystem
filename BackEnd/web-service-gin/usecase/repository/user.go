package repository

import (
	domain "github.com/luke92/FriendsLessonsSystem/domain"
)

type userDatabase struct {
	users []domain.User
}

func NewUserRepository() domain.UserRepository {
	lesson1 := domain.Lesson{Id: "1", Name: "Lesson 1", Type: "Math"}
	lesson2 := domain.Lesson{Id: "2", Name: "Lesson 2", Type: "Math"}
	lesson3 := domain.Lesson{Id: "3", Name: "Lesson 3", Type: "Math"}
	lesson4 := domain.Lesson{Id: "4", Name: "Lesson 4", Type: "Spanish"}
	lesson5 := domain.Lesson{Id: "5", Name: "Lesson 5", Type: "Spanish"}

	user2 := domain.User{Id: "2", Username: "mark", Password: "markPass", Email: "mark@gmail.com", Name: "Mark", Lastname: "Kennedy"}
	user3 := domain.User{Id: "3", Username: "jody", Password: "jodyPass", Email: "jody@gmail.com", Name: "Jody", Lastname: "Wagner"}
	user4 := domain.User{Id: "4", Username: "rachel", Password: "rachelPass", Email: "rachel@gmail.com", Name: "Rachel", Lastname: "Miles", Lessons: []domain.Lesson{lesson1, lesson2, lesson3, lesson4, lesson5}}
	user1 := domain.User{Id: "1", Username: "joe", Password: "joepass", Email: "joe@gmail.com", Name: "Joe", Lastname: "Reid", Friends: []domain.User{user4}}
	user5 := domain.User{Id: "5", Username: "luke", Password: "lukepass", Email: "luke@gmail.com", Name: "Luke", Lastname: "Vargas"}

	var users = []domain.User{
		user1,
		user2,
		user3,
		user4,
		user5,
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
