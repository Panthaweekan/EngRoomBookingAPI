package middlewares

import (
	"github.com/Panthaweekan/EngRoomBookingAPI/pkg/errors"
	"github.com/Panthaweekan/EngRoomBookingAPI/pkg/lodash"
	"github.com/gofiber/fiber/v2"
)

func AccessControlMiddleware(roleLimit []string) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		forbidden := errors.NewForbiddenError(errors.AuthErr("forbidden").Error())
		// unauth := errors.NewUnauthorizedError(errors.AuthErr("unauthorized").Error())
		access := true

		if !access {
			return lodash.ResponseError(c, forbidden)
		}

		return c.Next()
	}
}
