package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/yash0000001/p2psharingbackend/internal/database"
	"github.com/yash0000001/p2psharingbackend/internal/models"
	"github.com/yash0000001/p2psharingbackend/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/api/idtoken"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var body struct {
		Email    string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid Request body", err)
		return
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	user := models.User{
		Email:        body.Email,
		Username:     body.Username,
		PasswordHash: string(hash),
		CreatedAt:    time.Now(),
	}
	collection := database.DB.Collection("users")

	var existingUser models.User
	err := (collection.FindOne(
		context.Background(),
		bson.M{"email": body.Email},
	).Decode(&existingUser))

	if err == nil {
		utils.SendError(w, 409, "User already exists with this email", existingUser)
		return
	}
	err = (collection.FindOne(
		context.Background(),
		bson.M{"username": body.Username},
	).Decode(&existingUser))

	if err == nil {
		utils.SendError(w, 409, "This username has already been taken", existingUser)
		return
	}

	res, err := collection.InsertOne(context.Background(), user)

	if err != nil {
		utils.SendError(w, http.StatusBadGateway, "Bad Gateway Request", err)
		return
	}

	id := res.InsertedID.(primitive.ObjectID)

	frontendURL := os.Getenv("FRONTEND_URL")
	if frontendURL == "" {
		frontendURL = "http://localhost:3000"
	}
	frontendURL = strings.TrimRight(frontendURL, "/")
	verificationToken, _ := utils.GenerateVerificationToken(id.Hex())
	verificationLink := frontendURL + "/verify-email/token/" + verificationToken

	textBody := "Welcome to Blcak! Verify your email here: " + verificationLink
	htmlBody := utils.WelcomeEmailTemplate(user.Username, verificationLink)

	if err := utils.Mailer("Welcome to Blcak - Verify your email", user.Username, user.Email, textBody, htmlBody); err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to send verification email", err)
		return
	}

	utils.SendSuccess(w, http.StatusOK, "Signup successful. Check your email to verify your account.", map[string]interface{}{
		"user": map[string]interface{}{
			"id":       user.ID.Hex(),
			"email":    user.Email,
			"username": user.Username,
		},
	})
}

func VerifyEmail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var body struct {
		Token string `json:"token"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	userID, err := utils.VerifyVerificationToken(body.Token)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid verification token", err)
		return
	}

	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid user ID", err)
		return
	}

	collection := database.DB.Collection("users")
	_, err = collection.UpdateOne(
		context.Background(),
		bson.M{"_id": objID},
		bson.M{"$set": bson.M{"isVerified": true}},
	)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to verify email", err)
		return
	}

	utils.SendSuccess(w, http.StatusOK, "Email verified successfully", nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	collection := database.DB.Collection("users")

	var user models.User

	err := collection.FindOne(
		context.Background(),
		bson.M{"email": body.Email},
	).Decode(&user)

	if err != nil {
		utils.SendError(w, http.StatusNotFound, "User not found", err)
		return
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(body.Password),
	)

	if err != nil {
		utils.SendError(w, http.StatusUnauthorized, "Invalid Password", err)
		return
	}
	if !user.IsVerified || user.IsVerified == false {
		utils.SendError(w, http.StatusUnauthorized, "Please Verify Your Email", err)
		return
	}

	token, _ := utils.GenerateToken(user.ID.Hex())

	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    token,
		HttpOnly: true,
		Path:     "/",
		MaxAge:   3600 * 24,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
	})

	utils.SendSuccess(w, http.StatusOK, "User login successfull", map[string]interface{}{
		"user": map[string]interface{}{
			"id":       user.ID.Hex(),
			"email":    user.Email,
			"username": user.Username,
		},
	})
}

func GoogleSignin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var body struct {
		IDToken string `json:"idToken"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid Request body", err)
		return
	}

	payload, err := idtoken.Validate(context.Background(), body.IDToken, os.Getenv("GOOGLE_CLIENT_ID"))

	if err != nil {
		utils.SendError(w, http.StatusUnauthorized, "Invalid token", err)
		return
	}

	email := payload.Claims["email"].(string)

	collection := database.DB.Collection("users")

	var user models.User

	err = collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)

	if err != nil {

		user = models.User{
			Email:     email,
			Username:  payload.Claims["name"].(string),
			CreatedAt: time.Now(),
		}

		res, _ := collection.InsertOne(context.Background(), user)
		user.ID = res.InsertedID.(primitive.ObjectID)
	}

	token, _ := utils.GenerateToken(user.ID.Hex())

	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}

func ForgotPassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var body struct {
		Email string `json:"email"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	collection := database.DB.Collection("users")

	var user models.User

	err := collection.FindOne(context.Background(), bson.M{"email": body.Email}).Decode(&user)

	if err != nil {
		utils.SendError(w, http.StatusNotFound, "User not found", err)
		return
	}

	token := uuid.New().String()

	reset := models.PasswordReset{
		UserID:    user.ID,
		Token:     token,
		ExpiresAt: time.Now().Add(time.Minute * 15),
	}
	resetCollection := database.DB.Collection("password_resets")

	_, err = resetCollection.InsertOne(context.Background(), reset)
	if err != nil {
		http.Error(w, "Could not create reset token", 500)
		return
	}
	resetLink := "http://localhost:3000/reset-password?token=" + token

	text := "Reset your password using the link: " + resetLink

	html := utils.ResetPasswordTemplate(user.Username, resetLink)

	utils.Mailer(
		"Reset your password",
		user.Username,
		user.Email,
		text,
		html,
	)

	utils.SendSuccess(w, http.StatusOK, "Reset Email Sent", nil)
}

func ResetPassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var body struct {
		Token       string `json:"token"`
		NewPassword string `json:"newPassword"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid request body", 400)
		return
	}

	resetCollection := database.DB.Collection("password_resets")

	var reset models.PasswordReset

	err := resetCollection.FindOne(context.Background(), bson.M{"token": body.Token}).Decode(&reset)

	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	if time.Now().After(reset.ExpiresAt) {
		utils.SendError(w, http.StatusBadRequest, "Token Expired", 400)
		return
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(body.NewPassword), 10)

	userCollection := database.DB.Collection("users")

	userCollection.UpdateOne(context.Background(),
		bson.M{"_id": reset.UserID},
		bson.M{"$set": bson.M{"password_hash": string(hash)}},
	)

	resetCollection.DeleteOne(context.Background(), bson.M{"token": body.Token})

	utils.SendSuccess(w, http.StatusOK, "Password Reset Successfully", nil)
}

func Logout(w http.ResponseWriter, r *http.Request) {

	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   true,                  
		SameSite: http.SameSiteNoneMode, 
	})

	utils.SendSuccess(w, http.StatusOK, "User logged out successfully", nil)
}

func Me(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userID, ok := r.Context().Value("userID").(string)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	collection := database.DB.Collection("users")

	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		http.Error(w, "Invalid user ID", 400)
		return
	}

	var user models.User
	err = collection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		http.Error(w, "User not found", 404)
		return
	}

	utils.SendSuccess(w, 200, "User fetched", map[string]interface{}{
		"user": map[string]interface{}{
			"id":       user.ID.Hex(),
			"email":    user.Email,
			"username": user.Username,
		},
	})
}
