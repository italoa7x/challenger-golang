package domain

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Base struct {
	ID        string `json:"user_id"; gorm:type:"uuid;primary_key"`
	CreatedAt string `json: "created_at"; gorm:type:"datetime"`
	UpdatedAt string `json: "updated_at"; gorm:type:"datetime`
	DeletedAt string `json: "deleted_at; gorm:type:"datetime" sql:"index"`
}

func newUser() *User {
	return &User{}
}

func (base *Base) BeforeCreate(scope *gorm.Scope) error {
	err := scope.SetColumn("CreatedAt", time.Now())

	if err != nil {
		log.Fatalf("Error during obj creation, %v: ", err)
	}
	// cria um ID unico para o usuario, sempre que a estrutura Base for iniciada.
	idUser, errorIDUser := uuid.NewV4()
	// Atribiu para o usuario o ID gerado
	scope.SetColumn("ID", idUser.String())
	// Verifica se ocorreu  um erro na geracao do ID, se sim, retorna uma mesagem de erro
	if err != nil {
		log.Fatalf("Erro during obj ID, %v: ", err)
	}
	return nil
}
