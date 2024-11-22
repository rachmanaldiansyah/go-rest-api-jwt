package routes

import (
	"github.com/gorilla/mux"

	usercontroller "go-rest-api-jwt/controllers/user.controller"
	"go-rest-api-jwt/middleware"
)

func UserRouter(r *mux.Router) {
	router := r.PathPrefix("/users").Subrouter()
	router.Use(middleware.Auth)

	router.HandleFunc("/me", usercontroller.Me).Methods("GET")
}
