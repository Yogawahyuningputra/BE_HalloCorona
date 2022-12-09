package models

import "time"

type Consultation struct {
	ID            int       `json:"id" gorm:"primary_key:auto_increment"`
	Fullname      string    `json:"fullname" gorm:"type: varchar(255)"`
	Phone         string    `json:"phone" gorm:"type: varchar(255)"`
	BornDate      string    `json:"born_date" gorm:"type: varchar(255)"`
	Age           string    `json:"age" gorm:"type: varchar (255)"`
	Height        string    `json:"height" gorm:"type: varchar(255)"`
	Weight        string    `json:"weight" gorm:"type: varchar(255)"`
	Subject       string    `json:"subject" gorm:"type: varchar(255)"`
	ConsultanDate time.Time `json:"-"`
	Description   string    `json:"description" gorm:"type: varchar(255)"`
	UserID        int       `json:"user_id"`
	User          User      `json:"user"`
	CreateAt      time.Time `json:"-"`
	UpdateAt      time.Time `json:"-"`
}

type ConsultationResponse struct {
	ID       int    `json:"id"`
	Fullname string `json:"username"`
	Email    string `json:"email" gorm:"type: varchar(255)"`
}

func (ConsultationResponse) TableName() string {
	return "consultation"
}
