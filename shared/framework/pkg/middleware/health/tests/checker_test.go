package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/kittichai/core-framework/shared/framework/pkg/middleware/health"
)

func TestHealthCheck(t *testing.T) {

	healthChecker := health.NewHealthChecker(
		health.NewAPIChecker(),
	)
	app := fiber.New()

	app.Get("/health", func(c *fiber.Ctx) error {
		status := healthChecker.Check(c.Context())
		if status.Status == "healthy" {
			return c.JSON(status)
		}
		return c.Status(fiber.StatusServiceUnavailable).JSON(status)
	})

	t.Run("Healthy Status", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/health", nil)
		resp, err := app.Test(req)
		if err != nil {
			t.Fatalf("failed to make request: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("expected status 200, got %d", resp.StatusCode)
		}

		var status health.HealthStatus
		json.NewDecoder(resp.Body).Decode(&status)
		if status.Status != "healthy" {
			t.Fatalf("expected status 'healthy', got '%s'", status.Status)
		}
		if status.Details["api"] != "ok" {
			t.Fatalf("unexpected details: %v", status.Details)
		}
	})

}
