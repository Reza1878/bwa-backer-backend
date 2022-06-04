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

func (s *serviceImpl) GetTransactionsByUserId(userId int) ([]Transaction, error) {
	transactions, err := s.repository.FindByUserId(userId)

	return transactions, err
}

func (s *serviceImpl) CreateTransaction(request CreateTransactionRequest) (Transaction, error) {
	transaction := Transaction{}
	transaction.Amount = request.Amount
	transaction.CampaignId = request.CampaignId
	transaction.UserId = request.User.Id
	transaction.Status = "pending"

	newTransaction, err := s.repository.Save(transaction)

	return newTransaction, err
}
