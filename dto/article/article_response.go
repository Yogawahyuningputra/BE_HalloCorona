package articledto

import "be_corona/models"

type ArticleResponse struct {
	ID    int                 `json:"id" gorm:"primary_key:auto_increment"`
	Title string              `json:"title" gorm:"title"`
	Image string              `json:"image"`
	Desc  string              `json:"desc" gorm:"desc"`
	User  models.UserResponse `json:"user" gorm:"foreignKey:UserID"`
}
type ArticleDeleteResponse struct {
	ID    int    `json:"id"`
	Title string `json:"title" form:"title"`
}
