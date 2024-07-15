package validations

type LoginInput struct {
	Email    string `json:"email" binding:"required" type:"email"`
	Password string `json:"password" binding:"required"`
}
type RegisterInput struct {
	Fullname       string `json:"fullname" binding:"required"`
	Email          string `json:"email" binding:"required" type:"email"`
	Username       string `json:"username" binding:"required"`
	Password       string `json:"password" binding:"required"`
	ConfirmPasword string `json:"confirm_password" binding:"required"`
}
type CreateUserInput struct {
	Email    string `json:"email" binding:"required" type:"email"`
	Fullname string `json:"fullname" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
