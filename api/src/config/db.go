package config

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

func DB() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error Loading .env File")
	}

	DBMS := "mysql"
	USER := os.Getenv("USER")
	PASS := os.Getenv("PASS")
	DBNAME := os.Getenv("DBNM")
	PROTOCOL := os.Getenv("PTCL")

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic(err.Error())
	}

	return db
}
