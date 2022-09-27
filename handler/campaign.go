package handler

import (
	"bwa-backer/campaign"
	"bwa-backer/helper"
	"bwa-backer/user"
	"fmt"
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
		helper.ResponseBadRequest(c, "Get campaign failed", nil)
		return
	}
	helper.ResponseOK(c, "Get campaign success", campaign.FormatCampaigns(campaigns))
}

func (h *campaignHandler) GetCampaign(c *gin.Context) {
	var input campaign.GetCampaignDetailRequest

	err := c.ShouldBindUri(&input)

	if err != nil {
		helper.ResponseBadRequest(c, "Failed to get campaign", nil)
		return
	}

	campaignDetail, err := h.campaignService.GetCampaign(input)

	if err != nil {
		helper.ResponseBadRequest(c, "Failed to get campaign", nil)
		return
	}
	helper.ResponseOK(c, "Success to get campaign", campaign.FormatCampaignDetail(campaignDetail))
}

func (h *campaignHandler) CreateCampaign(c *gin.Context) {
	var request campaign.CreateCampaignRequest
	err := c.ShouldBindJSON(&request)

	if err != nil {
		errors := helper.FormatValidationError(err)
		helper.ResponseUnprocessableEntity(c, "Failed to create campaign", gin.H{"errors": errors})
		return
	}

	user := c.MustGet("currentUser").(user.User)

	request.User = user

	newCampaign, err := h.campaignService.CreateCampaign(request)

	if err != nil {
		helper.ResponseUnprocessableEntity(c, "Failed to create campaign", gin.H{"errors": err.Error()})
		return
	}
	helper.ResponseOK(c, "Success to create campaign", campaign.FormatCampaign(newCampaign))
}

func (h *campaignHandler) UpdateCampaign(c *gin.Context) {
	var requestId campaign.GetCampaignDetailRequest

	err := c.ShouldBindUri(&requestId)
	if err != nil {
		helper.ResponseBadRequest(c, "Failed to update campaign", nil)
		return
	}

	var requestData campaign.CreateCampaignRequest
	err = c.ShouldBindJSON(&requestData)

	if err != nil {
		errors := helper.FormatValidationError(err)
		helper.ResponseUnprocessableEntity(c, "Failed to update campaign", gin.H{"errors": errors})
		return
	}

	user := c.MustGet("currentUser").(user.User)

	requestData.User = user

	updatedCampaign, err := h.campaignService.UpdateCampaign(requestId, requestData)

	if err != nil {
		helper.ResponseUnprocessableEntity(c, "Failed to update campaign", gin.H{"errors": err.Error()})
		return
	}

	helper.ResponseOK(c, "Success to update campaign", campaign.FormatCampaignDetail(updatedCampaign))
}

func (h *campaignHandler) UploadImage(c *gin.Context) {

	var request campaign.CreateCampaignImageRequest
	err := c.ShouldBind(&request)
	if err != nil {
		helper.ResponseBadRequest(c, "Failed to upload campaign image", gin.H{"errors": helper.FormatValidationError(err)})
		return
	}

	file, err := c.FormFile("image")
	if err != nil {
		helper.ResponseUnprocessableEntity(c, "Failed to upload campaign image", gin.H{"is_uploaded": false})
		return
	}

	extension := filepath.Ext(file.Filename)
	path := fmt.Sprintf("images/campaign/%d-%s%s", request.CampaignID, uuid.New().String(), extension)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		helper.ResponseBadRequest(
			c,
			"Failed to upload campaign image",
			gin.H{
				"is_uploaded": false,
				"errors":      err.Error(),
			})
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	request.User = currentUser

	_, err = h.campaignService.CreateCampaignImage(request, path)

	if err != nil {
		os.Remove(path)
		helper.ResponseBadRequest(c, "Failed to upload campaign image", gin.H{"is_uploaded": false})
		return
	}
	helper.ResponseOK(c, "Success to upload campaign image", gin.H{"is_uploaded": true})
}
