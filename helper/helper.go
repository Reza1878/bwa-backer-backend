package helper

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

func APIResponse(message string, code int, status string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	return Response{
		Meta: meta,
		Data: data,
	}
}

func FormatValidationError(err error) []string {
	var errors []string
	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		errors = append(errors, err.Error())
		return errors
	}
	for _, e := range validationErrors {
		errors = append(errors, e.Error())
	}
	return errors
}

func GetDotEnvVariable(key string) string {
	godotenv.Load(".env")

	return os.Getenv(key)
}

func GetConnectionString() string {
	dbName := GetDotEnvVariable("DB_NAME")
	dbPass := GetDotEnvVariable("DB_PASS")
	dbUser := GetDotEnvVariable("DB_USER")
	dbHost := GetDotEnvVariable("DB_HOST")
	dbPort := GetDotEnvVariable("DB_PORT")

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbHost, dbUser, dbPass, dbName, dbPort)

	// return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
}

func ResponseOK(c *gin.Context, message string, data interface{}) {
	response := APIResponse(message, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}

func ResponseBadRequest(c *gin.Context, message string, data interface{}) {
	response := APIResponse(message, http.StatusBadRequest, "error", data)
	c.JSON(http.StatusBadRequest, response)
}

func ResponseNotFound(c *gin.Context, message string, data interface{}) {
	response := APIResponse(message, http.StatusNotFound, "error", data)
	c.JSON(http.StatusNotFound, response)
}

func ResponseUnprocessableEntity(c *gin.Context, message string, data interface{}) {
	response := APIResponse(message, http.StatusUnprocessableEntity, "error", data)
	c.JSON(http.StatusUnprocessableEntity, response)
}

func ResponseInternalServerError(c *gin.Context, message string, data interface{}) {
	response := APIResponse(message, http.StatusInternalServerError, "error", data)
	c.JSON(http.StatusInternalServerError, response)
}

func ResponseUnAuthorized(c *gin.Context, message string, data interface{}) {
	response := APIResponse(message, http.StatusUnauthorized, "error", data)
	c.JSON(http.StatusUnauthorized, response)
}

func JoinProjectPath(dir string) string {
	currentDir, _ := os.Getwd()

	projectPath := GetDotEnvVariable("PROJECT_PATH")

	if projectPath != "" {
		return filepath.Join(projectPath, "/", dir)
	}

	return filepath.Join(currentDir, "/", dir)
}
