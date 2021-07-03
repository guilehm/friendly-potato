package controllers

import (
	"fmt"
	"goapi/db"
	"log"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

var userCollection = db.OpenCollection("user")
var validate = validator.New()

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	ok := true
	msg := ""

	if err != nil {
		ok = false
		msg = fmt.Sprintf("password is incorrect")
	}

	return ok, msg
}
