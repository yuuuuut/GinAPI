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

//
var firebaseAuth *auth.Client

// GetFirebase は*auth.Clientを返します。
func GetFirebase() *auth.Client {
	return firebaseAuth
}

// InitFirebase はFirebaseAdminと接続します。
func InitFirebase() {
	opt := option.WithCredentialsFile(os.Getenv("FIREBASE_ADMIN_SDK_FILENAME"))

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic(err.Error())
	}

	auth, err := app.Auth(context.Background())
	if err != nil {
		panic(err.Error())
	}

	firebaseAuth = auth
}

// InitTestFirebase は InitFirebaseのTest用。
func InitTestFirebase() {
	opt := option.WithCredentialsFile(os.Getenv("FIREBASE_ADMIN_SDK_TEST_PATH"))

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic(err.Error())
	}

	auth, err := app.Auth(context.Background())
	if err != nil {
		panic(err.Error())
	}

	firebaseAuth = auth
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
