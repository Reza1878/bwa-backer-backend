package handler

import (
	"bwa-backer/auth"
	"bwa-backer/helper"
	"bwa-backer/user"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *UserHandler {
	return &UserHandler{
		userService: userService,
		authService: authService,
	}
}

func (handler *UserHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput
	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		helper.ResponseUnprocessableEntity(c, "Register account failed", errorMessage)
		return
	}

	isEmailAvailable, _ := handler.userService.IsEmailAvailable(user.CheckEmailInput{Email: input.Email})

	if !isEmailAvailable {
		errorMessage := gin.H{"errors": "Register account failed"}
		helper.ResponseBadRequest(c, "Email is already taken", errorMessage)
		return
	}

	newUser, err := handler.userService.RegisterUser(input)
	if err != nil {
		helper.ResponseBadRequest(c, "Register account failed", nil)
		return
	}

	token, err := handler.authService.GenerateToken(newUser.Id)
	if err != nil {
		helper.ResponseBadRequest(c, "Register account failed", nil)
		return
	}
	formatter := user.FormatUser(newUser, token)
	helper.ResponseOK(c, "Account has been registered", formatter)
}

func (h *UserHandler) Update(c *gin.Context) {
	var input user.UpdateUserInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		helper.ResponseUnprocessableEntity(c, "Update failed", gin.H{"errors": errors})
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	updatedUser, err := h.userService.UpdateUser(input)

	if err != nil {
		helper.ResponseInternalServerError(c, "Update failed", gin.H{"errors": err.Error()})
		return
	}
	helper.ResponseOK(c, "Update success!", user.FormatUser(updatedUser, ""))
}

func (h *UserHandler) UpdatePassword(c *gin.Context) {
	var input user.UpdateUserPasswordInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		helper.ResponseUnprocessableEntity(c, "Update password failed", gin.H{"errors": errors})
		return
	}
	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	err = h.userService.UpdatePassword(input)
	if err != nil {
		helper.ResponseBadRequest(c, "Failed to update password", gin.H{"errors": err.Error()})
		return
	}
	helper.ResponseOK(c, "Success update password", nil)
}

func (h *UserHandler) Logout(c *gin.Context) {
	var input user.LogoutInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		helper.ResponseBadRequest(c, "Logout failed", gin.H{"errors": errors})
		return
	}

	err = h.authService.DeleteRefreshToken(input.RefreshToken)
	if err != nil {
		helper.ResponseBadRequest(c, "Logout failed", gin.H{"errors": err.Error()})
		return
	}

	helper.ResponseOK(c, "Logout succecss", nil)
}

func (handler *UserHandler) Login(c *gin.Context) {
	var input user.LoginInput
	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		helper.ResponseUnprocessableEntity(c, "Login failed", errorMessage)
		return
	}

	loggedInUser, err := handler.userService.Login(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		helper.ResponseUnprocessableEntity(c, "Login failed", errorMessage)
		return
	}

	token, err := handler.authService.GenerateToken(loggedInUser.Id)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		helper.ResponseUnprocessableEntity(c, "Login failed", errorMessage)
		return
	}

	refreshToken, err := handler.authService.GenerateRefreshToken(loggedInUser.Id)

	if err != nil {
		helper.ResponseInternalServerError(c, "Failed to generate token", nil)
		return
	}
	helper.ResponseOK(c, "Login success", gin.H{"access_token": token, "refresh_token": refreshToken})
}

func (handler *UserHandler) CheckEmailAvailability(c *gin.Context) {
	var input user.CheckEmailInput
	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		helper.ResponseUnprocessableEntity(c, "Invalid request", errorMessage)
		return
	}

	isAvailable, err := handler.userService.IsEmailAvailable(user.CheckEmailInput{Email: input.Email})

	if err != nil {
		helper.ResponseUnprocessableEntity(c, "Internal server error", map[string]any{
			"errors": err.Error(),
		})
		return
	}

	message := "Email available"

	if !isAvailable {
		message = "Email not available"
	}

	helper.ResponseOK(c, message, gin.H{"is_available": isAvailable})
}

func (handler *UserHandler) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")
	if err != nil {
		helper.ResponseBadRequest(c, "Failed to upload avatar", gin.H{"is_uploaded": false})
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)

	path := fmt.Sprintf("images/avatar/%d-%s", currentUser.Id, file.Filename)

	err = c.SaveUploadedFile(file, path)

	if err != nil {
		helper.ResponseBadRequest(c, "Failed to upload avatar", gin.H{"is_uploaded": false})
		return
	}

	_, err = handler.userService.SaveAvatar(currentUser.Id, path)

	if err != nil {
		helper.ResponseBadRequest(c, "Failed to upload avatar", gin.H{"is_uploaded": false})
		return
	}
	if currentUser.Avatar != "" {
		os.Remove(currentUser.Avatar)
	}
	helper.ResponseOK(c, "Avatar uploaded successfully", gin.H{"is_uploaded": true})
}

func (h *UserHandler) FetchUser(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	formattedUser := user.FormatUser(currentUser, "")

	helper.ResponseOK(c, "Successfully fetch user data", formattedUser)
}
