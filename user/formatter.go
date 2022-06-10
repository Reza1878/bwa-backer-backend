package user

type UserResponse struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
	Email      string `json:"email"`
	Token      string `json:"token"`
	ImageUrl   string `json:"image_url"`
}

func FormatUser(user User, token string) UserResponse {
	formatter := UserResponse{
		Id:         user.Id,
		Name:       user.Name,
		Occupation: user.Occupation,
		Email:      user.Email,
		Token:      token,
		ImageUrl:   user.Avatar,
	}

	return formatter
}
