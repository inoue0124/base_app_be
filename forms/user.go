package forms

type UserSignup struct {
	Email    string `json:"email" binding:"required"`
	UserName string `json:"userName"`
}
