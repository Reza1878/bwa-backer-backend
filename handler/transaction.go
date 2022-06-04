package handler

import (
	"bwa-backer/helper"
	"bwa-backer/transaction"
	"bwa-backer/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	transactionService transaction.Service
}

func NewTransactionHandler(transactionService transaction.Service) *transactionHandler {
	return &transactionHandler{
		transactionService: transactionService,
	}
}

func (h *transactionHandler) GetTransactions(c *gin.Context) {
	var request transaction.GetCampaignTransactionRequest
	err := c.ShouldBindUri(&request)

	if err != nil {
		response := helper.APIResponse("Failed to get transactions", http.StatusUnprocessableEntity, "error", gin.H{
			"errors": helper.FormatValidationError(err),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	request.User = currentUser

	transactions, err := h.transactionService.GetTransactionsByCampaignId(request)

	if err != nil {
		response := helper.APIResponse("Failed to get transactions", http.StatusUnprocessableEntity, "error", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Success to get transactions", http.StatusOK, "success", transaction.FormatCampaignTransactions(transactions))
	c.JSON(http.StatusOK, response)
}

func (h *transactionHandler) GetUserTransaction(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)

	transactions, err := h.transactionService.GetTransactionsByUserId(currentUser.Id)

	if err != nil {
		response := helper.APIResponse("Failed to get transactions", http.StatusInternalServerError, "error", gin.H{"errors": "Internal server error"})
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helper.APIResponse("User's transactions", http.StatusOK, "success", transaction.FormatUserTransactions(transactions))
	c.JSON(http.StatusOK, response)
}
