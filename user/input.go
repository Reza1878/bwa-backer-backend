package user

type RegisterUserInput struct {
	Name       string `json:"name" binding:"required"`
	Occupation string `json:"occupation" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
}

type UpdateUserInput struct {
	Name       string `json:"name" binding:"required"`
	Occupation string `json:"occupation" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	User       User
}

type UpdateUserPasswordInput struct {
	OldPassword          string `json:"oldPassword" binding:"required" validate:"min=8"`
	NewPassword          string `json:"newPassword" binding:"required" validate:"min=8"`
	ConfirmationPassword string `json:"confirmationPassword" binding:"required" validate:"min=8"`
	User                 User
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LogoutInput struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type CheckEmailInput struct {
	Email string `json:"email" binding:"required,email"`
}
