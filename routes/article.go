package routes

import (
	"be_corona/handlers"
	"be_corona/pkg/middleware"
	"be_corona/pkg/mysql"
	"be_corona/repositories"

	"github.com/gorilla/mux"
)

func ArticleRoutes(r *mux.Router) {
	articleRepository := repositories.RepositoryArticle(mysql.DB)
	h := handlers.HandlerArticle(articleRepository)

	r.HandleFunc("/article", middleware.Auth(middleware.UploadFile(h.AddArticle))).Methods("POST")
	r.HandleFunc("/articles", middleware.Auth(h.FindArticles)).Methods("GET")
	r.HandleFunc("/article/{id}", h.GetArticle).Methods("GET")
	r.HandleFunc("/article/{id}", h.DeleteArticle).Methods("DELETE")

}
