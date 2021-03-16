package main

import (
	"github.com/yuuuuut/gin-api/src/db"
	"github.com/yuuuuut/gin-api/src/router"
)

func main() {
	db.DB()
	router.Init()
}
