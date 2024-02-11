package api

import (
	"20-HotelReservation/types"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func getAuthUser(c *fiber.Ctx) (*types.User, error) {
	user, ok := c.Context().UserValue("user").(*types.User)
	if !ok {
		return nil, fmt.Errorf("unauthorized")
	}
	return user, nil
}
