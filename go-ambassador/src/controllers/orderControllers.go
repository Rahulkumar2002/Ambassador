package controllers

import (
	"github.com/Rahulkumar2002/ambassador/src/database"
	"github.com/Rahulkumar2002/ambassador/src/models"
	"github.com/gofiber/fiber/v2"
)

func Order(c *fiber.Ctx) error {
	var orders []models.Order

	database.DB.Preload("OrderItems").Find(&orders)

	for i, order := range orders {
		orders[i].Name = order.FullName()
		orders[i].Total = order.GetTotal()
	}

	return c.JSON(orders)
}
