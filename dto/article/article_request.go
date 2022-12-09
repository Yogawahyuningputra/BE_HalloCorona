package articledto

type ArticleRequest struct {
	Title  string `json:"title" form:"title" validate:"required"`
	Image  string `json:"image" form:"image"`
	Desc   string `json:"desc" form:"desc" validate:"required"`
	UserID int    `json:"user_id"`
}
