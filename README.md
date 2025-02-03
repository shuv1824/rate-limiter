# High Performance Rate Limiter Service in Golang (API Gateway Component)
## Core Features
### Multiple Rate Limiting Strategies
- Fixed window
- Sliding window
- Token bucket
- Leaky bucket
### In-Memory & Distributed Storage
- Redis for distributed rate limiting
- Local in-memory storage for fast access
### Middleware for API Protection
- gRPC and REST API support
- Easy integration with any service
### Real-time Metrics & Monitoring
- Track blocked and allowed requests
- Expose Prometheus metrics
### Horizontal Scalability
- Support multiple instances with Redis/etcd as a shared store

## Golang Concepts
- **Goroutines & Channels** - Handle concurrent requests efficiently
- **Mutex & Atomic Counters** - Prevent race conditions in in-memory mode
- **Redis/etcd Integration** - Support distributed rate limiting
- **Middleware Pattern** - Make it pluggable for APIs
- **REST & gRPC** - Support multiple protocols
- **Testing & Benchmarking** - Ensure high performance under load

## Project Architecture
### Hexagonal Architecture Overview
- **Domain Layer** - Core logic (rate limiting strategies)
- **Application Layer** - Orchestrates rate limiting logic
- **Infrastructure Layer** - Handles storage, API, monitoring
