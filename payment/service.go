package payment

import (
	"bwa-backer/transaction"
	"bwa-backer/user"
)

type Service interface {
	GetToken(transaction transaction.Transaction, user user.User) (string, error)
}
