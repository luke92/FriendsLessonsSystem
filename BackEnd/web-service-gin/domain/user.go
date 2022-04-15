package domain

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
}

type UserRepository interface {
	GetAll() ([]User, error)
	GetByID(id string) (User, error)
}
