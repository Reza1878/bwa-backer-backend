package user

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	UpdateUser(input UpdateUserInput) (User, error)
	UpdatePassword(input UpdateUserPasswordInput) error
	Login(input LoginInput) (User, error)
	IsEmailAvailable(input CheckEmailInput) (bool, error)
	SaveAvatar(ID int, fileLocation string) (User, error)
	GetUserByID(ID int) (User, error)
}
