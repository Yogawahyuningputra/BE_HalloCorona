package routes

import (
	"be_corona/handlers"
	"be_corona/pkg/middleware"
	"be_corona/pkg/mysql"
	"be_corona/repositories"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
	userRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerAuth(userRepository)

	r.HandleFunc("/signup", h.Signup).Methods("POST")
	r.HandleFunc("/signin", h.Signin).Methods("POST")
	r.HandleFunc("/check-auth", middleware.Auth(h.CheckAuth)).Methods("GET")

}
