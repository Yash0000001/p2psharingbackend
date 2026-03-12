package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserStats struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID             primitive.ObjectID `bson:"user_id" json:"userId"`
	TotalSentBytes     int64              `bson:"total_sent_bytes" json:"totalSentBytes"`
	TotalReceivedBytes int64              `bson:"total_received_bytes" json:"totalReceivedBytes"`
	FilesSent          int64              `bson:"files_sent" json:"filesSent"`
	FilesReceived      int64              `bson:"files_received" json:"filesReceived"`
	UpdatedAt          time.Time          `bson:"updated_at" json:"updatedAt"`
}
