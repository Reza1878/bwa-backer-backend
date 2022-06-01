package transaction

import (
	"bwa-backer/campaign"
	"errors"
)

type serviceImpl struct {
	repository         Repository
	campaignRepository campaign.Repository
}

func NewService(repository Repository, campaignRepository campaign.Repository) *serviceImpl {
	return &serviceImpl{
		repository:         repository,
		campaignRepository: campaignRepository,
	}
}

func (s *serviceImpl) GetTransactionsByCampaignId(request GetCampaignTransactionRequest) ([]Transaction, error) {
	campaign, err := s.campaignRepository.FindByID(request.ID)
	if err != nil {
		return []Transaction{}, errors.New("cannot get campaign id")
	}
	if campaign.UserId != request.User.Id {
		return []Transaction{}, errors.New("not permitted")
	}
	transactions, err := s.repository.FindByCampaignId(request.ID)

	return transactions, err
}