package payment

import (
	"bwa-backer/helper"
	"bwa-backer/user"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type serviceImpl struct{}

func NewService() *serviceImpl {
	return &serviceImpl{}
}

func (s *serviceImpl) GetPaymentURL(transaction Transaction, user user.User) (string, error) {
	midtrans.ServerKey = helper.GetDotEnvVariable("MIDTRANS_SERVER_KEY")
	midtrans.Environment = midtrans.Sandbox

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  transaction.OrderId,
			GrossAmt: int64(transaction.Amount),
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: user.Name,
			Email: user.Email,
		},
	}
	snapTokenResp, err := snap.CreateTransaction(req)

	if err != nil {
		return "", err
	}

	return snapTokenResp.RedirectURL, nil
}

func (s *serviceImpl) ProcessPayment(request TransactionNotificationRequest) error {
	// transaction_id := request.OrderID

	return nil
}
