package auth

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"reqiured"`
}

type LoginResponce struct {
	Token string `json:"token"`
}

type RegisterResponce struct {
	Token string `json:"token"`
}
