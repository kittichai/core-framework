// templates/service-template/internal/interfaces/middleware/security.go
package middleware

import "github.com/gofiber/fiber/v2"

func SecurityHeadersMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Set("X-Content-Type-Options", "nosniff")
		c.Set("X-Frame-Options", "DENY")
		c.Set("Content-Security-Policy", "default-src 'self'")
		c.Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
		c.Set("X-XSS-Protection", "1; mode=block")
		return c.Next()
	}
}
