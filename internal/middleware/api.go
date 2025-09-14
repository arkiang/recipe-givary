package middleware

import "github.com/gofiber/fiber/v2"

func ForceJSON() fiber.Handler {
	return func(c *fiber.Ctx) error {
		method := c.Method()
		if method == fiber.MethodPost || method == fiber.MethodPut || method == fiber.MethodPatch {
			ct := c.Get("Content-Type")
			if ct == "" || ct == "text/plain" {
				// Force JSON as default
				c.Request().Header.SetContentType(fiber.MIMEApplicationJSON)
			}
		}
		return c.Next()
	}
}
