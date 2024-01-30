package controller

import (
	"encoding/json"
	database "go-base-fs/db"
	middlewares "go-base-fs/handlers"
	user_model "go-base-fs/models"
	"net/http"
)

var client = database.DatabaseConnection()
var Auths = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	validToken, err := middlewares.GenerateJWT()
	if err != nil {
		middlewares.ErrorResponse("Error generating token", response)
	}
	middlewares.SuccessResponse(validToken, response)
})

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user user_model.User
	decoder := json.NewDecoder(r.Body)
	// collection := client.Database(os.Getenv("DB_NAME")).Collection("USER")

	// collection.InsertOne(context.Background(), )
	if err := decoder.Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	// json.NewEncoder(w).Encode()
}
