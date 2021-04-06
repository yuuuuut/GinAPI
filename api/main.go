package main

import (
	"github.com/yuuuuut/gin-api/src/router"
	"github.com/yuuuuut/gin-api/src/util"
)

func main() {
	util.InitENV()
	//db := util.InitTestDB()
	util.InitDB()

	//util.InitCreateTables(db)
	util.InitFirebase()

	router.Init()

}
