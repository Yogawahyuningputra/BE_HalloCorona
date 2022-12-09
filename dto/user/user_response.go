package userdto

type UserResponse struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullname"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	ListAs   string `json:"list_as"`
	Gender   string `json:"gender"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
}
