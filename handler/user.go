package handler

import (
	"bwa-backer/auth"
	"bwa-backer/helper"
	"bwa-backer/user"
	"fmt"
	"net/http"

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
		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := handler.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "failed", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := handler.authService.GenerateToken(newUser.Id)
	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "failed", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := user.FormatUser(newUser, token)
	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (handler *UserHandler) Login(c *gin.Context) {
	var input user.LoginInput
	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedInUser, err := handler.userService.Login(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	token, err := handler.authService.GenerateToken(loggedInUser.Id)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := user.FormatUser(loggedInUser, token)
	response := helper.APIResponse("Login success", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (handler *UserHandler) CheckEmailAvailability(c *gin.Context) {
	var input user.CheckEmailInput
	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Invalid request", http.StatusUnprocessableEntity, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isAvailable, err := handler.userService.IsEmailAvailable(user.CheckEmailInput{Email: input.Email})

	if err != nil {
		response := helper.APIResponse("Internal server error", http.StatusUnprocessableEntity, "failed", map[string]any{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	message := "Email available"

	if !isAvailable {
		message = "Email not available"
	}

	response := helper.APIResponse(message, http.StatusOK, "success", map[string]any{
		"is_available": isAvailable,
	})
	c.JSON(http.StatusOK, response)
}

func (handler *UserHandler) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")
	if err != nil {
		response := helper.APIResponse("Failed to upload avatar", http.StatusBadRequest, "error", map[string]any{
			"is_uploaded": false,
		})
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)

	path := fmt.Sprintf("images/avatar/%d-%s", currentUser.Id, file.Filename)

	err = c.SaveUploadedFile(file, path)

	if err != nil {
		response := helper.APIResponse("Failed to upload avatar", http.StatusBadRequest, "error", map[string]any{
			"is_uploaded": false,
		})
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = handler.userService.SaveAvatar(currentUser.Id, path)

	if err != nil {
		response := helper.APIResponse("Failed to upload avatar", http.StatusBadRequest, "error", map[string]any{
			"is_uploaded": false,
		})
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Avatar uploaded successfully", http.StatusOK, "success", map[string]any{
		"is_uploaded": true,
	})
	c.JSON(http.StatusOK, response)
}

func (h *UserHandler) FetchUser(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	formattedUser := user.FormatUser(currentUser, "")

	response := helper.APIResponse("Successfully fetch user data", http.StatusOK, "success", formattedUser)
	c.JSON(http.StatusOK, response)
}
