package utils

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var SECRET_KEY = os.Getenv("JWT_SECRET_KEY")

func LogRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func HandleApiErrors(w http.ResponseWriter, status int, message string) {
	if message == "" {
		message = http.StatusText(status)
	}

	jsonResponse, _ := json.Marshal(struct {
		Error string `json:"error"`
	}{message})
	w.WriteHeader(status)
	w.Write(jsonResponse)
}

type SignedDetails struct {
	Email string
	Uid   string
	jwt.StandardClaims
}

func GenerateTokens(email string, uid string) (signedToken string, signedRefreshToken string, err error) {
	claims := &SignedDetails{
		Email: email,
		Uid:   uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Minute * time.Duration(10)).Unix(),
		},
	}

	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(
		jwt.SigningMethodHS256, claims,
	).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Panic(err)
		return
	}

	refreshToken, err := jwt.NewWithClaims(
		jwt.SigningMethodHS256, refreshClaims,
	).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Panic(err)
		return
	}

	return token, refreshToken, err
}

func UpdateTokens(
	signedToken string,
	signedRefreshToken string,
	userId string,
	collection *mongo.Collection,
) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	dateUpdated, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updateObj := primitive.D{
		bson.E{Key: "token", Value: signedToken},
		bson.E{Key: "refresh_token", Value: signedRefreshToken},
		bson.E{Key: "updated_at", Value: dateUpdated},
	}

	upsert := true
	filter := bson.M{"id": userId}
	opt := options.UpdateOptions{
		Upsert: &upsert,
	}

	_, err := collection.UpdateOne(
		ctx,
		filter,
		bson.D{
			{Key: "$set", Value: updateObj},
		},
		&opt,
	)
	defer cancel()

	if err != nil {
		log.Panic(err)
		return
	}

}
