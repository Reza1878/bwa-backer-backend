package transaction

import (
	"bwa-backer/campaign"
	"bwa-backer/payment"

	"errors"
)

type serviceImpl struct {
	repository         Repository
	campaignRepository campaign.Repository
	paymentService     payment.Service
}

func NewService(repository Repository, campaignRepository campaign.Repository, paymentService payment.Service) *serviceImpl {
	return &serviceImpl{
		repository:         repository,
		campaignRepository: campaignRepository,
		paymentService:     paymentService,
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

func (s *serviceImpl) GetTransactionSummary(request GetTransactionSummaryRequest) ([]TransactionSummary, error) {
	transactionSummary, err := s.repository.FindTransactionSummary(request.DateStart, request.DateEnd, request.UserId)
	return transactionSummary, err
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

	if err != nil {
		return newTransaction, err
	}

	paymentTransaction := payment.Transaction{
		OrderId: newTransaction.Code,
		Amount:  newTransaction.Amount,
	}

	paymentUrl, err := s.paymentService.GetPaymentURL(paymentTransaction, request.User)

	if err != nil {
		return newTransaction, err
	}

	newTransaction.PaymentUrl = paymentUrl

	newTransaction, err = s.repository.Update(newTransaction)

	return newTransaction, err
}

func (s *serviceImpl) ProcessTransaction(request TransactionNotificationRequest) error {
	transaction, err := s.repository.FindByTransactionCode(request.OrderID)
	if err != nil {
		return err
	}

	if request.PaymentType == "credit_card" && request.TransactionStatus == "capture" && request.FraudStatus == "accept" {
		transaction.Status = "paid"
	} else if request.TransactionStatus == "settlement" {
		transaction.Status = "paid"
	} else if request.TransactionStatus == "deny" || request.TransactionStatus == "expire" || request.TransactionStatus == "cancel" {
		transaction.Status = "cancelled"
	}

	updatedTransaction, err := s.repository.Update(transaction)
	if err != nil {
		return err
	}

	campaign, err := s.campaignRepository.FindByID(updatedTransaction.CampaignId)
	if err != nil {
		return err
	}

	if updatedTransaction.Status == "paid" {
		campaign.BackerCount = campaign.BackerCount + 1
		campaign.CurrentAmount = campaign.CurrentAmount + updatedTransaction.Amount

		_, err := s.campaignRepository.Update(campaign)
		if err != nil {
			return err
		}
	}
	return nil
}
