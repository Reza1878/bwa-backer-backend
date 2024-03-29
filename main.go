package main

import (
	"log"

	"bwa-backer/auth"
	"bwa-backer/campaign"
	"bwa-backer/handler"
	"bwa-backer/helper"
	"bwa-backer/middleware"
	"bwa-backer/payment"
	"bwa-backer/transaction"
	"bwa-backer/user"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := helper.GetConnectionString()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	campaignRepository := campaign.NewRepository(db)
	transactionRepository := transaction.NewRepository(db)
	authenticationRepository := auth.NewRepository(db)

	userService := user.NewService(userRepository)
	campaignService := campaign.NewService(campaignRepository)
	authService := auth.NewService(authenticationRepository)
	paymentService := payment.NewService()
	transactionService := transaction.NewService(transactionRepository, campaignRepository, paymentService)

	userHandler := handler.NewUserHandler(userService, authService)
	campaignHandler := handler.NewCampaignHandler(campaignService)
	transactionHandler := handler.NewTransactionHandler(transactionService)
	router := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{"Authorization", "Content-Type"}
	router.Use(cors.New(corsConfig))
	router.Static("/images", helper.JoinProjectPath("images"))

	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.PUT("/users", middleware.AuthMiddleware(authService, userService), userHandler.Update)
	api.PUT("/users/password", middleware.AuthMiddleware(authService, userService), userHandler.UpdatePassword)
	api.POST("/sessions", userHandler.Login)
	api.DELETE("/sessions", middleware.AuthMiddleware(authService, userService), userHandler.Logout)
	api.PUT("/sessions", userHandler.RefreshToken)

	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", middleware.AuthMiddleware(authService, userService), middleware.LimitUploadFileSize(512*1000), userHandler.UploadAvatar)
	api.GET("/users/fetch", middleware.AuthMiddleware(authService, userService), userHandler.FetchUser)

	api.GET("/campaigns", campaignHandler.GetCampaigns)
	api.GET("/campaigns/:id", campaignHandler.GetCampaign)
	api.POST("/campaigns", middleware.AuthMiddleware(authService, userService), campaignHandler.CreateCampaign)
	api.PUT("/campaigns/:id", middleware.AuthMiddleware(authService, userService), campaignHandler.UpdateCampaign)
	api.POST("/campaigns-images", middleware.AuthMiddleware(authService, userService), middleware.LimitUploadFileSize(512*1000), campaignHandler.UploadImage)
	api.DELETE("/campaigns-images/:campaign_image_id", middleware.AuthMiddleware(authService, userService), campaignHandler.DeleteImage)

	api.GET("/campaigns/:id/transactions", middleware.AuthMiddleware(authService, userService), transactionHandler.GetTransactions)
	api.GET("/transactions", middleware.AuthMiddleware(authService, userService), transactionHandler.GetUserTransaction)
	api.POST("/transactions", middleware.AuthMiddleware(authService, userService), transactionHandler.CreateTransaction)
	api.POST("/transactions/notification", transactionHandler.GetNotification)
	api.GET("/transactions/summary", middleware.AuthMiddleware(authService, userService), transactionHandler.GetTransactionSummary)

	router.Run()
}
