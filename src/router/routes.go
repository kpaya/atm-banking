package router

import "github.com/gofiber/fiber/v2"

func initializeRoutes(router *fiber.App) {
	// Create a api/v1 group
	v1 := router.Group("/api/v1")

	// Create a account group from api/v1
	account := v1.Group("/account")

	// Create account router
	account.Post("/create", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Create Account",
		})
	})

	// Withdraw router
	account.Post("/withdraw", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Do Withdraw",
		})
	})

	// Deposit router
	account.Post("/deposit", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Do Deposit",
		})
	})

	// Transfer router
	account.Post("/transfer", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Do Transfer",
		})
	})

	// Check Balance router
	account.Post("/balance", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Check Balance",
		})
	})
}
