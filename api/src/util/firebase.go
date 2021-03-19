package util

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

// RequestBody はFirebaseにCustomTokenをRequestする際の構造体
type RequestBody struct {
	Token             string `json:"token"`
	ReturnSecureToken bool   `json:"returnSecureToken"`
}

// ResponseBody はFirebaseからのResponse構造体
type ResponseBody struct {
	IdToken string
}

// SetupFirebase はFirebaseAdminのAuthClientを返します。
func SetupFirebase() (*auth.Client, error) {
	if err := godotenv.Load(".env"); err != nil {
		return nil, err
	}

	opt := option.WithCredentialsFile(os.Getenv("FIREBASE_ADMIN_SDK_FILENAME"))

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, err
	}

	auth, err := app.Auth(context.Background())
	if err != nil {
		return nil, err
	}

	return auth, nil
}

// GetVerifyIDToken は 使用可能のTokenを返します。
func GetVerifyIDToken(auth *auth.Client) (string, error) {
	url := os.Getenv("FIREBASE_URL")
	uid := os.Getenv("FIREBASE_ADMIN_USER_UID")

	token, err := auth.CustomToken(context.Background(), uid)
	if err != nil {
		return "", err
	}

	data := new(RequestBody)
	data.Token = token
	data.ReturnSecureToken = true

	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	res, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var d ResponseBody

	if err := json.Unmarshal(body, &d); err != nil {
		return "", err
	}

	return d.IdToken, nil
}
