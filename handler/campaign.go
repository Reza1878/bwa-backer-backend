package handler

import (
	"bwa-backer/campaign"
	"bwa-backer/helper"
	"bwa-backer/user"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type campaignHandler struct {
	campaignService campaign.Service
}

func NewCampaignHandler(campaignService campaign.Service) *campaignHandler {
	return &campaignHandler{
		campaignService: campaignService,
	}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.campaignService.GetCampaigns(userId)

	if err != nil {
		response := helper.APIResponse("Get campaign failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Get campaign success", http.StatusOK, "success", campaign.FormatCampaigns(campaigns))
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) GetCampaign(c *gin.Context) {
	var input campaign.GetCampaignDetailRequest

	err := c.ShouldBindUri(&input)

	if err != nil {
		response := helper.APIResponse("Failed to get campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	campaignDetail, err := h.campaignService.GetCampaign(input)

	if err != nil {
		response := helper.APIResponse("Failed to get campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to get campaign", http.StatusOK, "success", campaign.FormatCampaignDetail(campaignDetail))
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) CreateCampaign(c *gin.Context) {
	var request campaign.CreateCampaignRequest
	err := c.ShouldBindJSON(&request)

	if err != nil {
		errors := helper.FormatValidationError(err)
		response := helper.APIResponse("Failed to create campaign", http.StatusUnprocessableEntity, "error", map[string]any{
			"errors": errors,
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	user := c.MustGet("currentUser").(user.User)

	request.User = user

	newCampaign, err := h.campaignService.CreateCampaign(request)

	if err != nil {
		response := helper.APIResponse("Failed to create campaign", http.StatusUnprocessableEntity, "error", map[string]any{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Success to create campaign", http.StatusOK, "success", campaign.FormatCampaign(newCampaign))

	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) UpdateCampaign(c *gin.Context) {
	var requestId campaign.GetCampaignDetailRequest

	err := c.ShouldBindUri(&requestId)
	if err != nil {
		response := helper.APIResponse("Failed to update campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var requestData campaign.CreateCampaignRequest
	err = c.ShouldBindJSON(&requestData)

	if err != nil {
		errors := helper.FormatValidationError(err)
		response := helper.APIResponse("Failed to update campaign", http.StatusUnprocessableEntity, "error", map[string]any{
			"errors": errors,
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	user := c.MustGet("currentUser").(user.User)

	requestData.User = user

	updatedCampaign, err := h.campaignService.UpdateCampaign(requestId, requestData)

	if err != nil {
		response := helper.APIResponse("Failed to update campaign", http.StatusUnprocessableEntity, "error", map[string]any{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Success to update campaign", http.StatusOK, "success", campaign.FormatCampaignDetail(updatedCampaign))
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) UploadImage(c *gin.Context) {

	var request campaign.CreateCampaignImageRequest
	err := c.ShouldBind(&request)
	if err != nil {
		response := helper.APIResponse("Failed to upload campaign image", http.StatusBadRequest, "error", gin.H{
			"errors": helper.FormatValidationError(err),
		})
		c.JSON(http.StatusBadRequest, response)
		return
	}

	file, err := c.FormFile("image")
	if err != nil {
		response := helper.APIResponse("Failed to upload campaign image", http.StatusUnprocessableEntity, "error", map[string]any{
			"is_uploaded": false,
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	extension := filepath.Ext(file.Filename)
	path := fmt.Sprintf("images/campaign/%d-%s%s", request.CampaignID, uuid.New().String(), extension)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		response := helper.APIResponse("Failed to upload campaign image", http.StatusBadRequest, "error", gin.H{
			"is_uploaded": false,
			"errors":      err.Error(),
		})
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	request.User = currentUser

	_, err = h.campaignService.CreateCampaignImage(request, path)

	if err != nil {
		os.Remove(path)
		response := helper.APIResponse("Failed to upload campaign image", http.StatusBadRequest, "error", gin.H{
			"is_uploaded": false,
		})
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to upload campaign image", http.StatusOK, "success", gin.H{
		"is_uploaded": true,
	})
	c.JSON(http.StatusOK, response)
}
