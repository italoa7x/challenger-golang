package main

import (
	"fmt"
	"log"

	"github.com/italoa7x/simple-project-golang/aplication/repositories"
	"github.com/italoa7x/simple-project-golang/domain"
	"github.com/italoa7x/simple-project-golang/framework/utils"
)

func main() {
	db := utils.ConnectDB()

	user := domain.User{
		Name: "maria",
		Email: "Maria@gmail.com",
		Password: "maria90",
	}
	user.Prepare()

	userRepo := repositories.UserRepositoryDB{Db: db}
	result, err := userRepo.Insert(&user)

	if err != nil {
		log.Fatalf("Error during insert user in database, %v: ", err)
	}

	fmt.Println(result)
}
