package repositories

import (
	"log"

	"github.com/italoa7x/simple-project-golang/domain"
	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	Insert(user *domain.User) (*domain.User, error)
}

type UserRepositoryDB struct {
	Db *gorm.DB
}

func (repo *UserRepositoryDB) Insert(user *domain.User) (*domain.User, error) {
	err := user.Prepare()
	if err != nil {
		log.Fatalf("Error during the user validation, %v: ", err)
	}
	// salva no banco a estrutura contendo os dados do usuario
	err = repo.Db.Create(user).Error
	// verifica se a insercao ocorreu com sucesso
	if err != nil {
		log.Fatalf("Error to persist user, %v: ", err)
		return user, err
	}
	return user, nil
}
