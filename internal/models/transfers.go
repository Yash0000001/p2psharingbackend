package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transfer struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	SenderID   primitive.ObjectID `bson:"sender_id" json:"senderId"`
	ReceiverID primitive.ObjectID `bson:"receiver_id" json:"receiverId"`
	RoomID     string             `bson:"room_id" json:"roomId"`
	Filename   string             `bson:"filename" json:"filename"`
	Filesize   string             `bson:"filesize" json:"filesize"`
	Status     string             `bson:"status" json:"status"`
	CreatedAt  time.Time          `bson:"created_at" json:"createdAt"`
}
