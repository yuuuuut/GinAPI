package util

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"

	"github.com/yuuuuut/gin-api/src/entities"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
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

	DB = db

	fmt.Println("DB Connect OK !")

	return db
}

func InitTestDB() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error Loading .env File")
	}

	DBMS := "mysql"
	USER := os.Getenv("USER")
	PASS := os.Getenv("PASS")
	DBNAME := os.Getenv("TESTDBNM")
	PROTOCOL := os.Getenv("TESTPTCL")

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic(err.Error())
	}

	DB = db

	fmt.Println("TestDB Connect OK !")

	return db
}

func InitCreateTables(db *gorm.DB) {
	db.CreateTable(&entities.User{})
	db.CreateTable(&entities.Todo{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
}

func DropTables(db *gorm.DB) {
	db.DropTable(entities.User{}, entities.Todo{})
}

func GetDB() *gorm.DB {
	return DB
}
