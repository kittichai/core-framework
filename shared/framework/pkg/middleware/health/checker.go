package health

import (
	"context"
)

// APIChecker สำหรับตรวจสอบสถานะ API
type APIChecker struct{}

func NewAPIChecker() *APIChecker {
	return &APIChecker{}
}

func (a *APIChecker) Name() string {
	return "api"
}

func (a *APIChecker) Check(_ context.Context) (string, error) {
	return "ok", nil
}
