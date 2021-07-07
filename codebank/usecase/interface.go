package usecase

import (
	"github.com/codeedu/codebank/domain"
	"github.com/codeedu/codebank/dto"
)

type IUseCaseTransaction interface {
	ProcessTransaction(transactionDto dto.Transaction) (domain.Transaction, error)
}
