package controller

import (
	"context"
	"encoding/json"
	"fmt"
	database "go-base-fs/db"
	middlewares "go-base-fs/handlers"
	user_model "go-base-fs/models"
	"go-base-fs/utils"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

var client = database.DatabaseConnection()
var collection = client.Database(middlewares.GetEnvVar("DB_NAME")).Collection("USER")

var Login = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	var user user_model.User
	err := json.NewDecoder(request.Body).Decode(&user)

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}
	filter := bson.M{"email": user.Email}
	var u bson.M
	collection.FindOne(context.TODO(), filter).Decode(&u)
	valid := utils.CheckPasswordHash(user.Password, u["password"].(string))
	fmt.Println("valid: ", u["_id"], valid)
	validToken, err := middlewares.GenerateJWT()
	if err != nil {
		middlewares.ErrorResponse("Error generating token", response)
	}
	middlewares.SuccessResponse(validToken, response)
})

var CreateUserHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var user user_model.User
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	user.Password = utils.HashPassword(user.Password)
	result, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		log.Fatal("Error creating User")
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
})
