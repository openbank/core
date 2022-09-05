package metrics

import (
	health "google.golang.org/grpc/health/grpc_health_v1"
)

// HealthCheckServer is the implementation of the gRPC health check service.
type HealthCheckServer struct {
	health.UnimplementedHealthServer
}
