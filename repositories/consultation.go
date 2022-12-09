package repositories

import (
	"be_corona/models"

	"gorm.io/gorm"
)

type ConsultationRepository interface {
	// FindConsultations() ([]models.Consultation, error)
	// GetConsultation(ID int) (models.Consultation, error)
	AddConsultation(consutation models.Consultation) (models.Consultation, error)
	// UpdateConsultation(consultation models.Consultation) (models.Consultation, error)
	// DeleteConsultation(consultation models.Consultation) (models.Consultation, error)
}

func RepositoryConsultation(db *gorm.DB) *repository {
	return &repository{db}
}

// func (r *repository) FindConsultations() ([]models.Consultation, error) {
// 	var consultations []models.Consultation
// 	err := r.db.Preload("User").Find(&consultations).Error

// 	return consultations, err
// }

// func (r *repository) GetConsultation(ID int) (models.Consultation, error) {
// 	var consultation models.Consultation
// 	err := r.db.Preload("User").First(&consultation, ID).Error

// 	return consultation, err
// }

func (r *repository) AddConsultation(consultation models.Consultation) (models.Consultation, error) {
	err := r.db.Create(&consultation).Error

	return consultation, err
}

// func (r *repository) UpdateConsultation(consultation models.Consultation) (models.Consultation, error) {
// 	err := r.db.Save(&consultation).Error

// 	return consultation, err
// }

// func (r *repository) DeleteConsultation(consultation models.Consultation) (models.Consultation, error) {
// 	err := r.db.Delete(&consultation).Error

// 	return consultation, err
// }
