package consultationdto

import "time"

type ConsultationRequest struct {
	Fullname      string    `json:"fullname" form:"fullname" validate:"required"`
	Phone         string    `json:"phone" form:"phone" validate:"required"`
	BornDate      string    `json:"born_date" form:"born_date"`
	Age           string    `json:"age" form:"age" validate:"required"`
	Height        string    `json:"height" form:"height" validate:"required"`
	Weight        string    `json:"weight" form:"weight" validate:"required"`
	Subject       string    `json:"subject" form:"subject" validate:"required"`
	ConsultanDate time.Time `json:"-"`
	Description   string    `json:"description" form:"description" validate:"required"`
	UserID        int       `json:"user_id"`
	CreateAt      time.Time `json:"-"`
	UpdateAt      time.Time `json:"-"`
}
type ConsultationResponse struct {
	Fullname      string    `json:"fullname" form:"fullname" validate:"required"`
	Phone         string    `json:"phone" form:"phone" validate:"required"`
	BornDate      string    `json:"born_date" form:"born_date" validate:"required"`
	Age           string    `json:"age" form:"age" validate:"required"`
	Height        string    `json:"height" form:"height" validate:"required"`
	Weight        string    `json:"weight" form:"weight" validate:"required"`
	Subject       string    `json:"subject" form:"subject" validate:"required"`
	ConsultanDate time.Time `json:"-"`
	Description   string    `json:"description" form:"description" validate:"required"`
	UserID        int       `json:"user_id"`
	CreateAt      time.Time `json:"-"`
	UpdateAt      time.Time `json:"-"`
}
