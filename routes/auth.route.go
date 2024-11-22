package routes

import (
	"github.com/gorilla/mux"

	authcontroller "go-rest-api-jwt/controllers/auth.controller"
)

func AuthRouter(r *mux.Router) {
	router := r.PathPrefix("/auth").Subrouter()

	router.HandleFunc("/register", authcontroller.Register).Methods("POST")
}
