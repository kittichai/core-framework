// shared/infrastructure/tracing/otel.go
package tracing

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

func InitTracer(serviceName string) (trace.Tracer, error) {
	// Configure OpenTelemetry with Jaeger/Zipkin exporter
	tracer := otel.Tracer(serviceName)
	return tracer, nil
}
