package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Device struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID     primitive.ObjectID `bson:"user_id" json:"userId"`
	DeviceName string             `bson:"device_name" json:"deviceName"`
	DeviceType string             `bson:"device_type" json:"deviceType"`
	LastActive time.Time          `bson:"last_active" json:"lastActive"`
}
