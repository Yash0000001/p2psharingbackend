package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PasswordReset struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID    primitive.ObjectID `bson:"user_id" json:"userId"`
	Token     string             `bson:"token" json:"token"`
	ExpiresAt time.Time          `bson:"expires_at" json:"expiresAt"`
}
