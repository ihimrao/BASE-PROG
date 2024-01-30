package controller

import (
	"context"
	"encoding/json"
	database "go-base-fs/db"
	middlewares "go-base-fs/handlers"
	user_model "go-base-fs/models"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

var client = database.DatabaseConnection()
var Login = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	var user user_model.User
	err := json.NewDecoder(request.Body).Decode(&user)

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}
	// filter := bson.M{"email": user.Email}

	validToken, err := middlewares.GenerateJWT()
	if err != nil {
		middlewares.ErrorResponse("Error generating token", response)
	}
	middlewares.SuccessResponse(validToken, response)
})

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user user_model.User
	decoder := json.NewDecoder(r.Body)
	collection := client.Database(middlewares.GetEnvVar("DB_NAME")).Collection("USER")

	if err := decoder.Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	result, err := collection.InsertOne(context.Background(), bson.M{"username": "ihimrao", "email": "ihimrao@gmail.com", "password": "123456"})
	if err != nil {
		log.Fatal("Error creating User")
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
