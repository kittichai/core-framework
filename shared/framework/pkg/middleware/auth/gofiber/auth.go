// templates/service-template/internal/interfaces/middleware/auth.go
package gofiber

import (
	_ "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(jwtSecret string) fiber.Handler {
	//return jwt.New(jwt.Config{
	//	SigningKey: []byte(jwtSecret),
	//	ErrorHandler: func(c *fiber.Ctx, err error) error {
	//		return c.Status(401).JSON(errors.NewAppError(401, "Unauthorized", err))
	//	},
	//	ContextKey: "user",
	//})

	return nil
}
