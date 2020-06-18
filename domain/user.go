package domain

import (
	"log"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Base
	Name     string `json:"name"; gorm: type:"varchar(255)"`
	Email    string `json:"email"; gorm: type:"varchar(255); unique_index"`
	Password string `json: "-"; gorm: type:"varchar(255)"`
	Token    string `json: "acess_token"; gorm: type:"varchar(255)"; unique_index`
}

func (user *User) Prepare() error {
	// cria um Hash da senha
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	// Se não houver uma nova instância do ERRO, significa que a criação do Hash ocorreu com sucesso.
	if err != nil {
		log.Fatalf("Error during the password generation, %v: ", err)
	}
	user.Password = string(password)
	// gera um Token para o usuario
	token, erroToken := uuid.NewV4()
	user.Token = token.String()

	errorValidateUser = user.validate()

	// verifica se houve algum erro na validacao dos dados do usuario
	if errorValidateUser != nil {
		log.Fatalf("Error during the user validation, %v", error)
		return errorValidateUser
	}

	return nil
}

func (user *User) validate() error {

	if len(user.Name) == 0 {
		error = nil
		log.Fatalf("Nome do usuário inválido ou sem caracteris, %v: ", error)
	}

	if len(user.Email) == 0 {
		error = nil
		log.Fatalf("E-mail obrigatório, %v: ", error)
	}

	return error
}
