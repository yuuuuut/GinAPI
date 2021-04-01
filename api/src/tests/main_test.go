package tests

import (
	"os"
	"testing"

	"github.com/yuuuuut/gin-api/src/util"
)

func TestMain(m *testing.M) {
	util.InitTestENV()
	util.InitTestDB()
	util.InitTestFirebase()

	code := m.Run()

	os.Exit(code)
}
