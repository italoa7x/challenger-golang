type User struct {
	Base
	Name string
	Email string
	Password string
	Token string
}

func createUser() *User {
	return *domain.User{}
}

