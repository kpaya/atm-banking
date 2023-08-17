package repository

import (
	customerAgg "github.com/kpaya/atm-banking/src/aggregate/customer/entity"
)

type IAccountRepository interface {
	Get() (*customerAgg.Account, error)
}
