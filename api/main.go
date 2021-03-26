package main

import (
	"github.com/yuuuuut/gin-api/src/router"
	"github.com/yuuuuut/gin-api/src/util"
)

func main() {
	//util.InitDB()
	util.InitTestDB()

	router.Init()
}
