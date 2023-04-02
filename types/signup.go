package types

type SignupInput struct {
	Contact  string `json:"contact" binding:"required,email"`
	Password string `json:"password" binding:"required,gte=8,lte=24"`
}
