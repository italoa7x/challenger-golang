package domain

import (
	"log"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Base	 `valid:"required"`
	Name     string `gorm:"type:varchar(255)" valid:"notnull"`
	Email    string `gorm:"type:varchar(255);unique_index" valid:"notnull, email"`
	Password string `json:"-" valid:"notnull"`
	Token    string `json:"token" gorm:"unique_index" valid:"notnull, uuid"`
}

// primeira funcao a ser executada quando a estrutura for lida
func init() {
	// todos os campos da estrutura por padrão, não poderam ser nulos.
	govalidator.SetFieldsRequiredByDefault(true)
}

// funcao que retorna uma nova estrutura
func NewStructUser(name string, email string, password string) (*User, error) {
	// seta os dados para o usuario
	user := &User{
		Name:     name,
		Email:    email,
		Password: password,
	}
	// chama o metodo responsavel por preparar os hash's
	err := user.Prepare()
	//verifica se teve erro na preparacao
	if err != nil {
		return nil, err
	}
	// retorna OK se tudo deu certo
	return user, nil
}

// funcao que prepara o Hash da senha e do Token do usuario
func (user *User) Prepare() error {
	// cria um Hash da senha
	password, errorPassword := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	// Se não houver uma nova instância do ERRO, significa que a criação do Hash ocorreu com sucesso.
	if errorPassword != nil {
		log.Fatalf("Error during the password generation, %v: ", errorPassword)
		return errorPassword
	}
	// atribui o novo hash ao usuario
	user.Password = string(password)
	// gera um Token para o usuario
	token, errorToken := uuid.NewV4()
	// verifica se houve erro na criacao do token
	if errorToken != nil {
		log.Fatalf("Erro during create token")
		return errorToken
	}
	user.Token = token.String()
	errorValidateUser := user.validate()
	// verifica se houve algum erro na validacao dos dados do usuario
	if errorValidateUser != nil {
		log.Fatalf("Error during the user validation, %v", errorValidateUser)
		return errorValidateUser
	}
	return nil
}

func (user *User) validate() error {
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		return err
	}

	return nil
}
