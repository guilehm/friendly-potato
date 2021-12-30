package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"goapi/db"
	"goapi/models"
	"goapi/utils"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var usersCollection = db.OpenCollection("users", "")
var validate = validator.New()

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(providedPassword))
	ok := true
	msg := ""

	if err != nil {
		ok = false
		msg = "Password is incorrect"
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

	count, err := usersCollection.CountDocuments(ctx, bson.M{"email": user.Email})
	defer cancel()
	if err != nil {
		utils.HandleApiErrors(w, http.StatusInternalServerError, "")
		return
	}

	if count > 0 {
		utils.HandleApiErrors(w, http.StatusBadRequest, "Email already in use")
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

	_, err = usersCollection.InsertOne(ctx, user)
	if err != nil {
		utils.HandleApiErrors(w, http.StatusInternalServerError, "Could not create user")
		return
	}

	response, _ := json.Marshal(struct {
		Id      string `json:"id"`
		Token   string `json:"token"`
		Refresh string `json:"refresh_token"`
	}{user.UserId, *user.Token, *user.RefreshToken})
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func Login(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var userLogin models.UserLogin
	if err := json.NewDecoder(r.Body).Decode(&userLogin); err != nil {
		utils.HandleApiErrors(w, http.StatusBadRequest, err.Error())
		return
	}

	var user models.User
	if err := usersCollection.FindOne(
		ctx, bson.M{"email": userLogin.Email},
	).Decode(&user); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			utils.HandleApiErrors(w, http.StatusBadRequest, "User not found")
			return
		}
		utils.HandleApiErrors(w, http.StatusInternalServerError, "")
		return
	}

	ok, msg := VerifyPassword(*user.Password, userLogin.Password)
	if !ok {
		utils.HandleApiErrors(w, http.StatusBadRequest, msg)
		return
	}

	token, refresh, err := utils.GenerateTokens(*user.Email, user.UserId)
	if err != nil {
		utils.HandleApiErrors(w, http.StatusInternalServerError, "")
		return
	}

	tokens, err := utils.UpdateTokens(token, refresh, user.UserId, usersCollection)
	if err != nil {
		utils.HandleApiErrors(w, http.StatusInternalServerError, "")
		return
	}

	response, _ := json.Marshal(struct {
		Id      string `json:"id"`
		Token   string `json:"token"`
		Refresh string `json:"refresh_token"`
	}{user.UserId, *tokens.Token, *tokens.RefreshToken})
	w.Write(response)
}

func ValidateToken(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var tokens models.Tokens
	if err := json.NewDecoder(r.Body).Decode(&tokens); err != nil {
		utils.HandleApiErrors(w, http.StatusBadRequest, err.Error())
		return
	}

	var user models.User
	if err := usersCollection.FindOne(
		ctx, bson.M{"refresh_token": tokens.RefreshToken},
	).Decode(&user); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			utils.HandleApiErrors(w, http.StatusBadRequest, "Invalid token")
			return
		}
		utils.HandleApiErrors(w, http.StatusInternalServerError, "")
		return
	}

	response, _ := json.Marshal(struct {
		Ok bool `json:"ok"`
	}{Ok: true})
	w.Write(response)

}

func RefreshToken(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var tokens models.Tokens
	if err := json.NewDecoder(r.Body).Decode(&tokens); err != nil {
		utils.HandleApiErrors(w, http.StatusBadRequest, err.Error())
		return
	}

	var user models.User
	if err := usersCollection.FindOne(
		ctx, bson.M{"refresh_token": tokens.RefreshToken},
	).Decode(&user); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			utils.HandleApiErrors(w, http.StatusBadRequest, "Invalid token")
			return
		}
		utils.HandleApiErrors(w, http.StatusInternalServerError, "")
		return
	}

	token, refresh, err := utils.GenerateTokens(*user.Email, user.UserId)
	if err != nil {
		utils.HandleApiErrors(w, http.StatusInternalServerError, "")
		return
	}

	tokens, err = utils.UpdateTokens(token, refresh, user.UserId, usersCollection)
	if err != nil {
		utils.HandleApiErrors(w, http.StatusInternalServerError, "")
		return
	}

	response, _ := json.Marshal(struct {
		Id      string `json:"id"`
		Token   string `json:"token"`
		Refresh string `json:"refresh_token"`
	}{user.UserId, *tokens.Token, *tokens.RefreshToken})
	w.Write(response)

}
