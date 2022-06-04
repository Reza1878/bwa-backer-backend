package payment

import (
	"bwa-backer/transaction"
	"bwa-backer/user"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type serviceImpl struct{}

func NewService() *serviceImpl {
	return &serviceImpl{}
}

func (s *serviceImpl) GetToken(transaction transaction.Transaction, user user.User) (string, error) {
	midtrans.ServerKey = "YOUR-SERVER-KEY"
	midtrans.Environment = midtrans.Sandbox

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  string(rune(transaction.ID)),
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
