package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/yash0000001/p2psharingbackend/internal/database"
	"github.com/yash0000001/p2psharingbackend/internal/models"
	"github.com/yash0000001/p2psharingbackend/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const UserIDKey = "userID"

func CreateRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userIDVal := r.Context().Value(UserIDKey)
	if userIDVal == nil {
		utils.SendError(w, 401, "Unauthorized", nil)
		return
	}
	userID := userIDVal.(string)

	var body struct {
		Name       string  `json:"name"`
		Latitude   float64 `json:"lat"`
		Longitude  float64 `json:"lon"`
		DeviceName string  `json:"deviceName"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid body", err)
		return
	}
	userObjID, _ := primitive.ObjectIDFromHex(userID)

	room := models.Room{
		RoomID:     utils.UUID(),
		HostID:     userObjID,
		DeviceName: body.DeviceName,

		Location: models.Location{
			Type:        "Point",
			Coordinates: []float64{body.Longitude, body.Latitude}, // ⚠️ lng, lat order
		},

		Participants: []primitive.ObjectID{userObjID},

		IsActive: true,

		CreatedAt:  time.Now(),
		LastActive: time.Now(),
		ExpiresAt:  time.Now().Add(5 * time.Minute),
	}
	collection := database.DB.Collection("rooms")

	res, err := collection.InsertOne(context.Background(), room)
	if err != nil {
		utils.SendError(w, 500, "Failed to create room", err)
	}
	utils.SendSuccess(w, 200, "Room created", res.InsertedID)
}

func GetNearbyRooms(w http.ResponseWriter, r *http.Request) {
	lat, _ := strconv.ParseFloat(r.URL.Query().Get("lat"), 64)
	lng, _ := strconv.ParseFloat(r.URL.Query().Get("lng"), 64)

	collection := database.DB.Collection("rooms")

	delta := 0.002 // ~200m

	filter := bson.M{
		"is_active": true,
		"location.coordinates.0": bson.M{
			"$gte": lng - delta,
			"$lte": lng + delta,
		},
		"location.coordinates.1": bson.M{
			"$gte": lat - delta,
			"$lte": lat + delta,
		},
	}

	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		utils.SendError(w, 500, "DB error", err)
		return
	}

	var rooms []models.Room
	if err := cursor.All(context.Background(), &rooms); err != nil {
		utils.SendError(w, 500, "Decode error", err)
		return
	}

	utils.SendSuccess(w, 200, "Nearby rooms", rooms)
}

func JoinRoom(w http.ResponseWriter, r *http.Request) {

	userIDVal := r.Context().Value(UserIDKey)
	if userIDVal == nil {
		utils.SendError(w, 401, "Unauthorized", nil)
		return
	}
	userID := userIDVal.(string)

	var body struct {
		RoomID string `json:"roomId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid Body", err)
	}
	userObjID, _ := primitive.ObjectIDFromHex(userID)

	collection := database.DB.Collection("rooms")

	res, err := collection.UpdateOne(
		context.Background(),
		bson.M{"room_id": body.RoomID},
		bson.M{
			"$addToSet": bson.M{"participants": userObjID},
			"$set":      bson.M{"last_active": time.Now()},
		},
	)
	if err != nil {
		utils.SendError(w, 500, "Database Bad Request", err)
	}
	if res.MatchedCount == 0 {
		utils.SendError(w, 404, "Room not found", nil)
		return
	}
	utils.SendSuccess(w, 200, "Joined room", res)
}

func LeaveRoom(w http.ResponseWriter, r *http.Request) {
	userIDVal := r.Context().Value(UserIDKey)
	if userIDVal == nil {
		utils.SendError(w, 401, "Unauthorized", nil)
		return
	}
	userID := userIDVal.(string)
	var body struct {
		RoomID string `json:"roomId"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid Body", err)
	}
	// roomID, _ := primitive.ObjectIDFromHex(body.RoomID)
	userObjID, _ := primitive.ObjectIDFromHex(userID)

	collection := database.DB.Collection("rooms")

	res, err := collection.UpdateOne(context.Background(),
		bson.M{"room_id": body.RoomID},
		bson.M{
			"$pull": bson.M{"participants": userObjID},
			"$set":  bson.M{"last_active": time.Now()},
		},
	)
	if err != nil {
		utils.SendError(w, 500, "Database Bad Request", err)
	}
	if res.MatchedCount == 0 {
		utils.SendError(w, 404, "Room not found", nil)
		return
	}
	utils.SendSuccess(w, 200, "Room Left", res)
}

func DeleteRoom(w http.ResponseWriter, r *http.Request) {

	userIDVal := r.Context().Value(UserIDKey)
	if userIDVal == nil {
		utils.SendError(w, 401, "Unauthorized", nil)
		return
	}
	userID := userIDVal.(string)

	roomID := r.URL.Query().Get("roomId")

	objID, _ := primitive.ObjectIDFromHex(roomID)
	log.Println(objID)
	collection := database.DB.Collection("rooms")

	var room models.Room
	collection.FindOne(context.Background(), bson.M{"room_id": roomID}).Decode(&room)

	if room.HostID.Hex() != userID {
		utils.SendError(w, 403, "Only host can delete", nil)
		return
	}

	collection.DeleteOne(context.Background(), bson.M{"room_id": roomID})

	utils.SendSuccess(w, 200, "Room deleted", nil)
}
