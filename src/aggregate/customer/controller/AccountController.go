package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/kpaya/atm-banking/src/aggregate/customer/repository"
	usecase "github.com/kpaya/atm-banking/src/aggregate/customer/usecase"
	infra "github.com/kpaya/atm-banking/src/infra/database"

	dto "github.com/kpaya/atm-banking/src/aggregate/customer/usecase/dto"
)

func AccessAccount(c *fiber.Ctx) error {
	inputAccessAccountDTO := new(dto.InputAccessAccountDTO)

	if err := c.BodyParser(&inputAccessAccountDTO); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": fmt.Sprintf("internal error: %s", err.Error()),
		})
	}

	// create repository
	repository := repository.NewAccountRepository(infra.NewDatabaseInstance())

	// // create usecase
	usecase := usecase.NewAccessAccountUseCase(repository)
	// call usecase
	output, err := usecase.Execute(inputAccessAccountDTO)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": fmt.Sprintf("internal error: %s", err.Error()),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&output)
}
