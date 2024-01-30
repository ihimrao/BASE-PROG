package route

import (
	controller "go-base-fs/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

// Routes -> define endpoints
func Routes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/createUser", controller.CreateUserHandler).Methods("POST")
	router.HandleFunc("/login", controller.Login).Methods("POST")
	router.HandleFunc("/ping", func(response http.ResponseWriter, request *http.Request) {
		response.Write([]byte("pong!"))
	}).Methods("GET")
	// router.HandleFunc("/people", middlewares.IsAuthorized(controller.CreateUserHandler())).Methods("GET")
	// router.HandleFunc("/person/{id}", controllers.GetPersonEndpoint).Methods("GET")
	// router.HandleFunc("/person/{id}", controllers.DeletePersonEndpoint).Methods("DELETE")
	// router.HandleFunc("/person/{id}", controllers.UpdatePersonEndpoint).Methods("PUT")
	// router.HandleFunc("/upload", controllers.UploadFileEndpoint).Methods("POST")
	// router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./uploaded/"))))
	return router
}
