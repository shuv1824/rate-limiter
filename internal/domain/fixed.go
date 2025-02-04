package domain

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type FixedWindowLimiter struct {
	mu       sync.Mutex
	requests map[string]int
	limit    int
	window   time.Duration
}

func NewFixedWindowLimiter(limit int, window time.Duration) *FixedWindowLimiter {
	limiter := &FixedWindowLimiter{
		requests: make(map[string]int),
		limit:    limit,
		window:   window,
	}

	go limiter.cleanup()

	return limiter
}

func (f *FixedWindowLimiter) Allow(key string) bool {
	f.mu.Lock()
	defer f.mu.Unlock()

	now := time.Now().Unix()
	windowKey := key + ":" + strconv.FormatInt(now/int64(f.window.Seconds()), 10)

	count, exists := f.requests[windowKey]

	if !exists {
		f.requests[windowKey] = 1
		return true
	}

	if count >= f.limit {
		return false
	}

	f.requests[windowKey]++
	return true
}

func (f *FixedWindowLimiter) cleanup() {
	ticker := time.NewTicker(f.window)
	defer ticker.Stop()

	for range ticker.C {
		f.mu.Lock()
		now := time.Now().Unix()
		for key := range f.requests {
			if isExpired(key, now, f.window) {
				delete(f.requests, key)
			}
		}
	}
}

func isExpired(key string, now int64, window time.Duration) bool {
	var timestamp int64
	_, err := fmt.Sscanf(key, "%*s:%d", &timestamp)
	if err != nil {
		return false
	}

	return now/int64(window.Seconds()) > timestamp
}
