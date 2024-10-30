package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func envConnection() {
	varErr := godotenv.Load()
	if varErr != nil {
		log.Fatal("Error loading .env file")
	}
}

func DBConnection() {

	envConnection()

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	passwd := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		user, passwd, host, port, dbname,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn))

	if err != nil {
		panic("Error al conectar a la base de datos: " + err.Error())
	}
}
