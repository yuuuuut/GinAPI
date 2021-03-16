package db

import (
	"fmt"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/yuuuuut/gin-api/src/entity"
)

var db *gorm.DB

func DB() {
	time.Sleep(time.Second * 5)

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
	db, err = gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic(err.Error())
	}
	_ = db

	fmt.Println("DB Connect OK !")
}

func InitMigration() {
	db.AutoMigrate(&entity.Todo{})
}

func GetDB() *gorm.DB {
	return db
}
