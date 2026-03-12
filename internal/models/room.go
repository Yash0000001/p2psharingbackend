package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Room struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	RoomID     string             `bson:"room_id" json:"roomId"`
	HostID     primitive.ObjectID `bson:"host_id" json:"hostId"`
	DeviceName string             `bson:"device_name" json:"deviceName"`

	Latitude  float64 `bson:"latitude" json:"latitude"`
	Longitude float64 `bson:"longitude" json:"longitude"`

	Participants []primitive.ObjectID `bson:"participants" json:"participants"`

	IsActive bool `bson:"is_active" json:"isActive"`

	CreatedAt  time.Time `bson:"created_at" json:"createdAt"`
	LastActive time.Time `bson:"last_active" json:"lastActive"`

	ExpiresAt time.Time `bson:"expires_at" json:"expiresAt"`
}
