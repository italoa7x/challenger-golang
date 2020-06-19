package utils

import (
	"log"

	"github.com/italoa7x/simple-project-golang/domain"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

// funcao responsavel por conectar ao banco de dados
func ConnectDB() *gorm.DB {
	dbInfo := "dbname=go sslmode=disable user=postgres password=postgres"

	db, errorConnect := gorm.Open("postgres", dbInfo)
	// verifica se a conexao foi efetuada com sucesso, caso contrario, retorna um erro e impede a aplicaçaõ de subir
	if errorConnect != nil {
		log.Fatalf("Error connectin to DB, %v: ", errorConnect)
		panic(errorConnect)
	}
	// Passa a estrutura de usuario para gerar a tabela no banco de dados
	db.AutoMigrate(&domain.User{})

	return db
}
