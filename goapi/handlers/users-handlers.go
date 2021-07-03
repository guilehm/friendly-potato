package handlers

import (
	"context"
	"encoding/json"
	"goapi/db"
	"goapi/models"
	"goapi/utils"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

var userCollection = db.OpenCollection("user")
var validate = validator.New()

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
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

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.HandleApiErrors(w, http.StatusBadRequest, err.Error())
		return
	}

	if validationErr := validate.Struct(user); validationErr != nil {
		utils.HandleApiErrors(w, http.StatusBadRequest, validationErr.Error())
		return
	}

	count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
	defer cancel()
	if err != nil {
		utils.HandleApiErrors(w, http.StatusInternalServerError, "")
		return
	}

	if count > 0 {
		utils.HandleApiErrors(w, http.StatusBadRequest, "Please choose another email")
		return
	}

	password, err := HashPassword(*user.Password)
	if err != nil {
		utils.HandleApiErrors(w, http.StatusInternalServerError, "")
		return
	}

	user.ID = primitive.NewObjectID()
	user.UserId = user.ID.Hex()
	user.Password = &password
	user.DateAdded, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.DateChanged, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	token, refresh, err := utils.GenerateTokens(*user.Email, user.UserId)
	if err != nil {
		utils.HandleApiErrors(w, http.StatusInternalServerError, "")
		return
	}
	user.Token = &token
	user.RefreshToken = &refresh

	result, err := userCollection.InsertOne(ctx, user)
	if err != nil {
		utils.HandleApiErrors(w, http.StatusInternalServerError, "Could not create user")
		return
	}

	response, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
