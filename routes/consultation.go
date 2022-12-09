package routes

import (
	"be_corona/handlers"
	"be_corona/pkg/mysql"
	"be_corona/repositories"

	"github.com/gorilla/mux"
)

func ConsultationRoutes(r *mux.Router) {
	consultationRepository := repositories.RepositoryConsultation(mysql.DB)
	h := handlers.HandlerConsultation(consultationRepository)

	r.HandleFunc("/consultation", h.AddConsultation).Methods("POST")

}
