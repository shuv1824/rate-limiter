package domain

type RateLimiter interface {
	Allow(key string) bool
}
