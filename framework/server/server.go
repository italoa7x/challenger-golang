package main

import (
	"github.com/italoa7x/simple-project-golang/aplication/repositories"
	"github.com/italoa7x/simple-project-golang/domain"
)

func main() {

	user := domain.User{
		Name:     "italo",
		Email:    "italo@gmail.com",
		Password: "italo123",
	}

	userRepo := repositories.UserRepositoryDB{}
}
