package userdto

type CreateUserRequest struct {
	Fullname string `json:"fullname" form:"fullname" validate:"required"`
	Username string `json:"username" form:"username" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	ListAs   string `json:"list_as" form:"list_as" validate:"required"`
	Gender   string `json:"gender" form:"gender" validate:"required"`
	Phone    string `json:"phone" form:"phone" validate:"required"`
	Address  string `json:"address" form:"address" validate:"required"`
}

type UpdateUserRequest struct {
	Fullname string `json:"fullname" form:"fullname"`
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	ListAs   string `json:"list_as" form:"list_as"`
	Gender   string `json:"gender" form:"gender"`
	Phone    string `json:"phone" form:"phone"`
	Address  string `json:"address" form:"address"`
}
