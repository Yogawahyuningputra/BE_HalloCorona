package handlers

import (
	authdto "be_corona/dto/auth"
	resultdto "be_corona/dto/result"
	"be_corona/models"
	"be_corona/pkg/bcrypt"
	jwtToken "be_corona/pkg/jwt"
	"be_corona/repositories"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
)

type handlerAuth struct {
	AuthRepository repositories.AuthRepository
}

func HandlerAuth(AuthRepository repositories.AuthRepository) *handlerAuth {
	return &handlerAuth{AuthRepository}
}

func (h *handlerAuth) Signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(authdto.AuthRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := resultdto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := resultdto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	password, err := bcrypt.HashingPassword(request.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := resultdto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	user := models.User{
		Fullname: request.Fullname,
		Username: request.Username,
		Email:    request.Email,
		Password: password,
		ListAs:   request.ListAs,
		Gender:   request.Gender,
		Phone:    request.Phone,
		Address:  request.Address,
	}

	data, err := h.AuthRepository.Signup(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := resultdto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := resultdto.SuccessResult{Code: http.StatusOK, Data: convertResponse(data)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerAuth) Signin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(authdto.SigninRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := resultdto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	user := models.User{
		Username: request.Username,
		Password: request.Password,
	}

	user, err := h.AuthRepository.Signin(user.Username)
	fmt.Println("user login", user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := resultdto.ErrorResult{Code: http.StatusBadRequest, Message: "Wrong email"}
		json.NewEncoder(w).Encode(response)
		return
	}

	isValid := bcrypt.CheckPasswordHash(request.Password, user.Password)
	if !isValid {
		w.WriteHeader(http.StatusBadRequest)
		response := resultdto.ErrorResult{Code: http.StatusBadRequest, Message: "Wrong email or password"}
		json.NewEncoder(w).Encode(response)
		return
	}

	claims := jwt.MapClaims{}
	claims["id"] = user.ID
	claims["list_as"] = user.ListAs
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix() // 2 jam expired

	token, errGenerateToken := jwtToken.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		fmt.Println("Unauthorize")
		return
	}

	SigninResponse := authdto.SigninResponse{
		Username: user.Username,
		Status:   user.ListAs,
		Token:    token,
	}

	w.Header().Set("Content-Type", "application/json")
	response := resultdto.SuccessResult{Code: http.StatusOK, Data: SigninResponse}
	json.NewEncoder(w).Encode(response)

}

func (h *handlerAuth) CheckAuth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	// Check User by Id
	user, err := h.AuthRepository.Getuser(userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := resultdto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	CheckAuthResponse := authdto.CheckAuthResponse{
		ID:       user.ID,
		Fullname: user.Fullname,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
		ListAs:   user.ListAs,
		Gender:   user.Gender,
		Phone:    user.Phone,
		Address:  user.Address,
	}

	w.Header().Set("Content-Type", "application/json")
	response := resultdto.SuccessResult{Code: http.StatusOK, Data: CheckAuthResponse}
	json.NewEncoder(w).Encode(response)
}
