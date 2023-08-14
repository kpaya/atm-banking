package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	dto "github.com/kpaya/atm-banking/src/aggregate/customer/usecase/dto"
)

func AccessAccount(c *fiber.Ctx) error {
	inputAccessAccountDTO := new(dto.InputAccessAccountDTO)

	if err := c.BodyParser(&inputAccessAccountDTO); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": fmt.Sprintf("internal error: %s", err.Error()),
		})
	}

	// call usecase
	// outputAccessAccountDTO, err := usecase.AccessAccount(inputAccessAccountDTO)

	return c.Status(fiber.StatusOK).JSON(inputAccessAccountDTO)
}
