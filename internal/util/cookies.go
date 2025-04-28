package util

import "github.com/gofiber/fiber/v2"

func SetCookie(c *fiber.Ctx, name string, value string, maxAge int) {
	c.Cookie(&fiber.Cookie{Name: name, Value: value, MaxAge: maxAge,
		HTTPOnly: true,
		Secure:   true,
		SameSite: "lax",
	})
}
