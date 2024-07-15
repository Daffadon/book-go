package validations

type CreateAuthorInput struct {
	Fullname string `json:"fullname" binding:"required"`
	Email    string `json:"email" binding:"required" type:"email"`
}
