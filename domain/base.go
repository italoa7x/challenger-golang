package domain

import (
	"log"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Base struct {
	ID string `json:"user_id"; gorm:type:"uuid;primary_key"`
	//CreatedAt time.Time `json: "created_at"`
	//UpdatedAt time.Time `json: "updated_at"`
	//DeletedAt time.Time `json: "deleted_at`
}

func newUser() *User {
	return &User{}
}

func (base *Base) BeforeCreate(scope *gorm.Scope) error {
	/*err := scope.SetColumn("CreatedAt", time.Now())

	if err != nil {
		log.Fatalf("Error during obj creation, %v: ", err)
	}
	*/
	// cria um ID unico para o usuario, sempre que a estrutura Base for iniciada.
	idUser, errorIDUser := uuid.NewV4()
	// Atribiu para o usuario o ID gerado
	if errorIDUser != nil {
		log.Fatalf("Error during create user ID")
	}
	scope.SetColumn("ID", idUser.String())

	return nil
}
