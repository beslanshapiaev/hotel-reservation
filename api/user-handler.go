package api

import (
	"github.com/beslanshapiaev/hotel-reservation/types"
	"github.com/gofiber/fiber/v2"
)

func HandleGetUsers(c *fiber.Ctx) error {
	user := types.User{
		FirstName: "Mike",
		LastName:  "Smith",
	}

	user2 := types.User{
		FirstName: "Antonny",
		LastName:  "Pettis",
	}

	return c.JSON([]types.User{user, user2})
}

func HandleGetUser(c *fiber.Ctx) error {
	user := types.User{
		FirstName: "Mike",
		LastName:  "Smith",
	}
	return c.JSON(user)
}
