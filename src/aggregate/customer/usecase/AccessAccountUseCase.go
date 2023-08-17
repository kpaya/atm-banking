package usecase

import (
	"github.com/google/uuid"
	repository "github.com/kpaya/atm-banking/src/aggregate/customer/repository"
	"github.com/kpaya/atm-banking/src/aggregate/customer/usecase/dto"
)

type AccessAccountUseCase struct {
	Repository repository.IAccountRepository
}

func NewAccessAccountUseCase(repository repository.IAccountRepository) *AccessAccountUseCase {
	return &AccessAccountUseCase{
		Repository: repository,
	}
}

func (a *AccessAccountUseCase) Execute(input *dto.InputAccessAccountDTO) (*dto.OutputAccessAccountDTO, error) {
	_, err := a.Repository.Get()
	if err != nil {
		return nil, err
	}
	return &dto.OutputAccessAccountDTO{
		Key: uuid.NewString(),
	}, nil
}
