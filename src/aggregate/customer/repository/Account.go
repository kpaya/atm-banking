package repository

import (
	entity "github.com/kpaya/atm-banking/src/aggregate/customer/entity"
	"gorm.io/gorm"
)

type AccountRepositoryInterface interface {
	Get() (*entity.Account, error)
}

type AccountRepository struct {
	Db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{
		Db: db,
	}
}

func (a *AccountRepository) Get() (*entity.Account, error) {
	return &entity.Account{}, nil
}
