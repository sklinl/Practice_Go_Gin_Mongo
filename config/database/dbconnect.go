package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB_CONNECTOR *gorm.DB
var DB_ERROR error

func CreateDBConnection() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:morale@tcp(127.0.0.1:3306)/TEST?charset=utf8mb4&parseTime=True&loc=Local"
	DB_CONNECTOR, DB_ERROR = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if DB_ERROR != nil {
		log.Fatal(DB_ERROR)
	}
}
