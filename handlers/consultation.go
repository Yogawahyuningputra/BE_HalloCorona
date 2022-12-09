package handlers

import (
	consultationdto "be_corona/dto/consultation"
	resultdto "be_corona/dto/result"
	"be_corona/models"
	"be_corona/repositories"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
)

type handlerConsultation struct {
	ConsultationRepository repositories.ConsultationRepository
}

func HandlerConsultation(ConsultationRepository repositories.ConsultationRepository) *handlerConsultation {
	return &handlerConsultation{ConsultationRepository}
}

func (h *handlerConsultation) AddConsultation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))
	request := new(consultationdto.ConsultationRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := resultdto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := resultdto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	consultation := models.Consultation{
		Fullname:      request.Fullname,
		Phone:         request.Phone,
		BornDate:      request.BornDate,
		Age:           request.Age,
		Height:        request.Height,
		Weight:        request.Weight,
		Subject:       request.Subject,
		ConsultanDate: request.ConsultanDate,
		Description:   request.Description,
		UserID:        userId,
		// User: models.UserResponse,
		// Status : "pending",

	}
	consultation, err = h.ConsultationRepository.AddConsultation(consultation)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := resultdto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	ConsultationResponse := consultationdto.ConsultationResponse{
		Fullname:      consultation.Fullname,
		Phone:         consultation.Phone,
		BornDate:      consultation.BornDate,
		Age:           consultation.Age,
		Height:        consultation.Height,
		Weight:        consultation.Weight,
		Subject:       consultation.Subject,
		ConsultanDate: consultation.ConsultanDate,
		Description:   consultation.Description,
		// UserID:        consultation.UserID,
	}
	w.WriteHeader(http.StatusOK)
	response := resultdto.SuccessResult{Code: http.StatusOK, Data: ConsultationResponse}
	json.NewEncoder(w).Encode(response)

}
