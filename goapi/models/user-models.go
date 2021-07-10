package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRefresh struct {
	Token        *string `bson:"token" json:"token"`
	RefreshToken *string `bson:"refresh_token" json:"refresh_token"`
}

type UserLogin struct {
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"required"`
}

type User struct {
	ID           primitive.ObjectID `bson:"_id"`
	UserId       string             `bson:"id" json:"id"`
	Email        *string            `bson:"email" json:"email" validate:"email,required"`
	Password     *string            `bson:"password" json:"password" validate:"required"`
	Token        *string            `bson:"token" json:"token"`
	RefreshToken *string            `bson:"refresh_token" json:"refresh_token"`
	DateAdded    time.Time          `bson:"date_added" json:"date_added"`
	DateChanged  time.Time          `bson:"date_changed" json:"date_changed"`
}
