package utils

import (
	"context"
	"goapi/models"
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SignedDetails struct {
	Email string
	Uid   string
	jwt.StandardClaims
}

const accessTokenLifetime = time.Minute * time.Duration(10)
const refreshTokenLifetime = time.Hour * time.Duration(24)

func GenerateTokens(email string, uid string) (signedToken string, signedRefreshToken string, err error) {
	claims := &SignedDetails{
		Email: email,
		Uid:   uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(accessTokenLifetime).Unix(),
		},
	}

	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(refreshTokenLifetime).Unix(),
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
) (models.Tokens, error) {

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
		return models.Tokens{}, err
	}

	updatedTokens := models.Tokens{
		Token:        &signedToken,
		RefreshToken: &signedRefreshToken,
	}
	return updatedTokens, nil

}

func ValidateToken(signedToken string) (claims *SignedDetails, msg string) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)

	if err != nil {
		msg = err.Error()
		return nil, msg
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = "the token is invalid"
		return nil, msg
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = "token is expired"
		return nil, msg
	}

	return claims, msg
}
