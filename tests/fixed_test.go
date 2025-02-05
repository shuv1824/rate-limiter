package tests

import (
	"testing"
	"time"

	"github.com/shuv1824/rate-limiter/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestFixedWindowLimiter(t *testing.T) {
	limiter := domain.NewFixedWindowLimiter(5, time.Second)
	key := "test_user"

	for i := 0; i < 5; i++ {
		assert.True(t, limiter.Allow(key), "Request should be allowed")
	}

	assert.False(t, limiter.Allow(key), "Request should be blocked")

	time.Sleep(time.Second)
	assert.True(t, limiter.Allow(key), "Request should be allowed after reset")
}
