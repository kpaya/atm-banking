package usecase

import (
	repository_dto "github.com/kpaya/atm-banking/src/aggregate/customer/repository"
	"github.com/kpaya/atm-banking/src/aggregate/customer/usecase/dto"
)

type AccessAccountUseCase struct {
	Repository *repository_dto.AccountRepositoryInterface
}

func NewAccessAccountUseCase(repository *repository_dto.AccountRepositoryInterface) *AccessAccountUseCase {
	return &AccessAccountUseCase{
		Repository: repository,
	}
}

func (a *AccessAccountUseCase) Execute(input *dto.InputAccessAccountDTO) (*dto.OutputAccessAccountDTO, error) {

	return &dto.OutputAccessAccountDTO{}, nil
}
