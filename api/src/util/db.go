package util

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/yuuuuut/gin-api/src/entities"
)

var DB *gorm.DB

// InitDB は MainDatabaseと接続して*gorm.DBを返します。
func InitDB() *gorm.DB {
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

// InitTestDB は TestDatabaseと接続して*gorm.DBを返します。
func InitTestDB() *gorm.DB {
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

// InitCreateTables は与えられたDBにTableを作成します。
func InitCreateTables(db *gorm.DB) {
	db.CreateTable(&entities.User{})
	db.CreateTable(&entities.Profile{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.CreateTable(&entities.Todo{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.CreateTable(&entities.Tag{}).AddForeignKey("parent_id", "tags(id)", "CASCADE", "CASCADE")
	db.CreateTable(&entities.Comment{}).AddForeignKey("todo_id", "todos(id)", "CASCADE", "CASCADE").AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE").AddForeignKey("parent_id", "comments(id)", "CASCADE", "CASCADE")
	db.Table("todos_tags").AddForeignKey("todo_id", "todos(id)", "CASCADE", "CASCADE")
	db.Table("todos_tags").AddForeignKey("tag_id", "tags(id)", "CASCADE", "CASCADE")
}

// DropTables は与えられたDBのTableを削除します。
func DropTables(db *gorm.DB) {
	db.DropTable(entities.User{}, entities.Profile{}, entities.Todo{}, &entities.Tag{})
}

// GetDB は*gorm.DBを返します。
func GetDB() *gorm.DB {
	return DB
}
