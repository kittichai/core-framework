package health

import (
	"context"
	"time"
)

// Checker interface กำหนด contract สำหรับ component checkers
type Checker interface {
	Check(ctx context.Context) (string, error)
	Name() string
}

// HealthStatus เก็บผลลัพธ์ของ health check
type HealthStatus struct {
	Status  string            `json:"status"` // "healthy" or "unhealthy"
	Details map[string]string `json:"details"`
}

// HealthChecker จัดการการตรวจสอบหลาย component
type HealthChecker struct {
	checkers []Checker
}

// NewHealthChecker สร้าง HealthChecker ด้วย list ของ checkers
func NewHealthChecker(checkers ...Checker) *HealthChecker {
	return &HealthChecker{checkers: checkers}
}

// Check ทำการตรวจสอบทุก component และคืน HealthStatus
func (h *HealthChecker) Check(ctx context.Context) HealthStatus {
	status := HealthStatus{
		Status:  "healthy",
		Details: make(map[string]string),
	}

	for _, checker := range h.checkers {
		checkCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
		defer cancel()

		result, err := checker.Check(checkCtx)
		if err != nil {
			status.Status = "unhealthy"
			status.Details[checker.Name()] = "failed: " + err.Error()
		} else {
			status.Details[checker.Name()] = result
		}
	}

	return status
}
