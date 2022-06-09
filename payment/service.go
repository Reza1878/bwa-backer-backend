package payment

import (
	"bwa-backer/user"
)

type Service interface {
	GetPaymentURL(transaction Transaction, user user.User) (string, error)
	ProcessPayment(request TransactionNotificationRequest) error
}
