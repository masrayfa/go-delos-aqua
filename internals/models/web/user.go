package web

type UserRead struct {
	UserId   int    `json:"id_user"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type UserCreate struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserUpdate struct {
	UserId   int    `json:"id_user"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}