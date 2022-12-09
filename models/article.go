package models

import "time"

type Article struct {
	ID       int          `json:"id" gorm:"primary_key:auto_increment"`
	UserID   int          `json:"user_id" `
	User     UserResponse `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Title    string       `json:"title" gorm:"type: varchar(255)"`
	Image    string       `json:"image" gorm:"type:varchar(255)" form:"image"`
	Desc     string       `json:"desc" gorm:"type:text" form:"desc"`
	CreateAt time.Time    `json:"-"`
	UpdateAt time.Time    `json:"-"`
}
type ArticleResponse struct {
	ID     int          `json:"id"`
	UserID int          `json:"user_id"`
	User   UserResponse `json:"user" `
	Title  string       `json:"title"`
	Image  string       `json:"image"`
	Desc   string       `json:"desc"`
}
