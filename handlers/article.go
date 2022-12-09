package handlers

import (
	articledto "be_corona/dto/article"
	resultdto "be_corona/dto/result"
	"be_corona/models"
	"be_corona/repositories"
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
)

type handlerArticle struct {
	ArticleRepository repositories.ArticleRepository
}

// var path_file = "http://localhost:5000/uploads/"

func HandlerArticle(ArticleRepository repositories.ArticleRepository) *handlerArticle {
	return &handlerArticle{ArticleRepository}
}

func (h *handlerArticle) AddArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	dataContex := r.Context().Value("dataFile") // add this code
	filename := dataContex.(string)

	request := articledto.ArticleRequest{
		Title: r.FormValue("title"),
		Desc:  r.FormValue("desc"),
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := resultdto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	article := models.Article{
		Title:  request.Title,
		Image:  filename,
		Desc:   request.Desc,
		UserID: userId,
	}

	article, err = h.ArticleRepository.AddArticle(article)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := resultdto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	article, _ = h.ArticleRepository.GetArticle(article.ID)
	article.Image = os.Getenv("PATH_FILE") + article.Image
	w.WriteHeader(http.StatusOK)
	response := resultdto.SuccessResult{Code: http.StatusOK, Data: article}
	json.NewEncoder(w).Encode(response)
}
func (h *handlerArticle) FindArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	articles, err := h.ArticleRepository.FindArticles()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := resultdto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Create Embed Path File on Image property here ...
	for i, p := range articles {
		articles[i].Image = os.Getenv("PATH_FILE") + p.Image
	}

	w.WriteHeader(http.StatusOK)
	response := resultdto.SuccessResult{Code: http.StatusOK, Data: articles}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerArticle) GetArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var article models.Article
	article, err := h.ArticleRepository.GetArticle(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := resultdto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Create Embed Path File on Image property here ...
	article.Image = os.Getenv("PATH_FILE") + article.Image

	w.WriteHeader(http.StatusOK)
	response := resultdto.SuccessResult{Code: http.StatusOK, Data: convertResponseArticle(article)}
	json.NewEncoder(w).Encode(response)
}
func convertResponseArticle(u models.Article) models.ArticleResponse {
	return models.ArticleResponse{
		Title: u.Title,
		Desc:  u.Desc,
		Image: u.Image,
		User:  u.User,
	}
}
func (h *handlerArticle) DeleteArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	article, err := h.ArticleRepository.GetArticle(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := resultdto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.ArticleRepository.DeleteArticle(article)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := resultdto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	ProductDeleteResponse := articledto.ArticleDeleteResponse{
		ID:    userId,
		Title: data.Title,
	}

	w.WriteHeader(http.StatusOK)
	response := resultdto.SuccessResult{Code: http.StatusOK, Data: ProductDeleteResponse}
	json.NewEncoder(w).Encode(response)

}
