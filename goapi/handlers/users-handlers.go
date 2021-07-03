package handlers

import (
	"encoding/json"
	"goapi/db"
	"goapi/models"
	"log"
	"net/http"

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
		msg = "password is incorrect"
	}

	return ok, msg
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		panic(err)
	}

	if validationErr := validate.Struct(user); validationErr != nil {
		jsonResponse, _ := json.Marshal(struct {
			Error string `json:"error"`
		}{validationErr.Error()})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonResponse)
		return
	}

	jsonResponse, _ := json.Marshal(user)
	w.Write(jsonResponse)
}
