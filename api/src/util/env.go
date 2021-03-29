package util

import "github.com/joho/godotenv"

// InitENV は godotenv.Loadします。
func InitENV() {
	if err := godotenv.Load(".env"); err != nil {
		panic("Error Loading .env File")
	}
}

// InitTestENV は InitENVのTest用
func InitTestENV() {
	if err := godotenv.Load("../../.env"); err != nil {
		panic("Error Loading .env File")
	}
}
