package utils

import (
	"log"
	"os"

	"github.com/italoa7x/simple-project-golang/domain"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConnectDB() *gorm.DB {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file, %v: ", err)
	}
	DB_INFO := os.Getenv("DB_INFO")

	db, errorConnect := gorm.Open("postgres", DB_INFO)
	// verifica se a conexao foi efetuada com sucesso, caso contrario, retorna um erro e impede a aplicaçaõ de subir
	if errorConnect != nil {
		log.Fatalf("Error connectin to DB, %v: ", errorConnect)
		panic(errorConnect)
	}
	// Passa a estrutura de usuario para gerar a tabela no banco de dados
	db.AutoMigrate(&domain.User{})

	return db
}
