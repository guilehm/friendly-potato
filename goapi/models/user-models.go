package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id"`
	UserId       string             `json:"id"`
	Email        *string            `json:"email" validate:"email,required"`
	Password     *string            `json:"password" validate:"required"`
	Token        *string            `json:"token"`
	RefreshToken *string            `json:"refresh_token"`
	DateAdded    time.Time          `json:"date_added"`
	DateChanged  time.Time          `json:"date_changed"`
}
