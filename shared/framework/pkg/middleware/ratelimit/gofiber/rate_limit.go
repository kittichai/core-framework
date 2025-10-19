// templates/service-template/internal/interfaces/middleware/rate_limit.go
package pkg

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/kittichai/core-framework/shared/framework/pkg/errors"
)

func RateLimitMiddleware(rate, capacity int) fiber.Handler {
	return limiter.New(limiter.Config{
		Max:        rate,
		Expiration: time.Second,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(429).JSON(errors.NewAppError(429, "Too many requests", nil))
		},
	})
}
