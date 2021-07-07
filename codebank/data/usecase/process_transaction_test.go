package usecase_test

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/codeedu/codebank/data/usecase"
	"github.com/codeedu/codebank/domain"
	"github.com/codeedu/codebank/domain/mocks"
	"github.com/codeedu/codebank/dto"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func newCreditCard() domain.CreditCard {
	cc := domain.NewCreditCard()
	cc.Number = "1234"
	cc.Name = "Wesley"
	cc.ExpirationYear = 2021
	cc.ExpirationMonth = 7
	cc.CVV = 123
	cc.Limit = 1000
	cc.Balance = 0

	return *cc
}
func Test_ProcessTransaction(t *testing.T) {
	a := dto.Transaction{
		ID:              uuid.NewV4().String(),
		Name:            "John Doe",
		Number:          "12345678",
		ExpirationMonth: 7,
		ExpirationYear:  1996,
		CVV:             123,
		Amount:          12.00,
		Store:           "Jane's store",
		Description:     "Jane's store description",
		CreatedAt:       time.Now(),
	}

	cc := newCreditCard()
	testCases := []struct {
		name          string
		arg           dto.Transaction
		builtSts      func(repo *mocks.TransactionRepository)
		checkResponse func(t *testing.T, transaction domain.Transaction, err error)
	}{
		{
			name: "fail on get credit card",
			arg:  a,
			builtSts: func(repo *mocks.TransactionRepository) {
				repo.On("GetCreditCard", mock.Anything).Return(domain.CreditCard{}, errors.New("Unexpected Error"))
			},
			checkResponse: func(t *testing.T, transaction domain.Transaction, err error) {
				assert.Error(t, err)
				assert.Empty(t, transaction)
			},
		},
		{
			name: "fail on save transaction",
			arg:  a,
			builtSts: func(repo *mocks.TransactionRepository) {
				repo.On("GetCreditCard", mock.Anything).Return(cc, nil)
				repo.On("SaveTransaction", mock.Anything, mock.Anything).Return(errors.New("Unexpected Error"))
			},
			checkResponse: func(t *testing.T, transaction domain.Transaction, err error) {
				assert.Error(t, err)
				assert.Empty(t, transaction)
			},
		},
		{
			name: "success",
			arg:  a,
			builtSts: func(repo *mocks.TransactionRepository) {
				repo.On("GetCreditCard", mock.Anything).Return(cc, nil)
				repo.On("SaveTransaction", mock.Anything, mock.Anything).Return(nil)
			},
			checkResponse: func(t *testing.T, transaction domain.Transaction, err error) {
				fmt.Println(cc)
				assert.NoError(t, err)
				assert.NotEmpty(t, transaction)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := new(mocks.TransactionRepository)
			tc.builtSts(repo)
			u := usecase.NewUseCaseTransaction(repo)
			res, err := u.ProcessTransaction(tc.arg)
			tc.checkResponse(t, res, err)
		})
	}
}
