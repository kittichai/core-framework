

go-micro-framework/
├── shared/                  // Shared modules สำหรับ import
│   ├── domain/              // Abstract interfaces (reuse ทั่ว services ถ้ามี domain ร่วม)
│   │   ├── repository.go    // Generic Repository interface
│   │   ├── cache.go         // Cache interface
│   │   ├── event_publisher.go // Event Publisher interface
│   │   └── config_store.go  // Config Store interface (e.g., etcd)
│   ├── infrastructure/      // Reusable adapters
│   │   ├── persistence/     // Database adapters
│   │   │   ├── postgres/
│   │   │   ├── mysql/
│   │   │   └── mongo/
│   │   ├── cache/
│   │   │   └── redis/
│   │   ├── messaging/
│   │   │   ├── kafka/
│   │   │   └── rabbitmq/
│   │   ├── config/
│   │   │   └── etcd/
│   │   └── tracing/         // OpenTelemetry for tracing
│   │       └── otel.go
│   ├── pkg/                 // Shared utilities
│   │   ├── config/          // Viper-based config loader
│   │   ├── logger/          // Zap logger with tracing
│   │   ├── errors/          // Custom error handling
│   │   └── di/              // Dependency Injection (wire.go)
│   └── go.mod               // Module: github.com/yourorg/framework/shared
├── templates/               // Boilerplate สำหรับ new service
│   ├── service-template/    // Clone เพื่อสร้าง service ใหม่
│   │   ├── cmd/
│   │   │   └── main.go      // Entry point with DI
│   │   ├── internal/
│   │   │   ├── domain/      // Service-specific entities/interfaces
│   │   │   ├── app/         // Use cases
│   │   │   ├── infrastructure/ // Custom adapters
│   │   │   └── interfaces/  // Handlers (HTTP/gRPC)
│   │   │       ├── http/
│   │   │       ├── grpc/
│   │   │       └── middleware/
│   │   ├── Dockerfile       // สำหรับ containerize
│   │   ├── docker-compose.yml // Local dev
│   │   ├── config.yaml      // Example config
│   │   └── go.mod           // Depend on shared module
├── examples/                // ตัวอย่าง services
│   └── user-service/        // Built from template
├── scripts/                 // Helpers
│   └── new-service.sh       // Script สำหรับ generate new service
├── README.md                // How to use framework
└── go.work                  // สำหรับ develop framework locally