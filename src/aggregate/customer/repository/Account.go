package repository

import (
	"github.com/google/uuid"
	entity "github.com/kpaya/atm-banking/src/aggregate/customer/entity"
	"gorm.io/gorm"
)

type AccountRepository struct {
	Db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) IAccountRepository {
	return &AccountRepository{
		Db: db,
	}
}

func (a *AccountRepository) Get() (*entity.Account, error) {

	return &entity.Account{
		ID:               uuid.NewString(),
		AccountNumber:    "####",
		TotalBalance:     1000,
		AvaliableBalance: 1000,
	}, nil
}
