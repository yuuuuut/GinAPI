package util

import (
	"github.com/go-playground/validator/v10"
)

func TodoValidation(err error) (result map[string]string) {
	result = make(map[string]string)

	errors := err.(validator.ValidationErrors)
	if len(errors) != 0 {
		for i := range errors {
			switch errors[i].StructField() {
			case "Title":
				switch errors[i].Tag() {
				case "required":
					result["Title"] = "Titleを入力してください。"
				}
			case "UserID":
				switch errors[i].Tag() {
				case "required":
					result["UserId"] = "ログインが必要です。"
				}
			}
		}
	}

	return result
}
