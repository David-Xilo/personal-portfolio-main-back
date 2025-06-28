package middleware

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"net/http"
	"sync"
)

type IPRateLimiter struct {
	ips   map[string]*rate.Limiter
	mu    *sync.RWMutex
	limit rate.Limit
	burst int
}

func NewIPRateLimiter(limit rate.Limit, burst int) *IPRateLimiter {
	return &IPRateLimiter{
		ips:   make(map[string]*rate.Limiter),
		mu:    &sync.RWMutex{},
		limit: limit,
		burst: burst,
	}
}

func (ipRateLimiter *IPRateLimiter) GetLimiter(ip string) *rate.Limiter {
	ipRateLimiter.mu.Lock()
	defer ipRateLimiter.mu.Unlock()

	limiter, exists := ipRateLimiter.ips[ip]
	if !exists {
		limiter = rate.NewLimiter(ipRateLimiter.limit, ipRateLimiter.burst)
		ipRateLimiter.ips[ip] = limiter
	}

	return limiter
}

func RateLimitMiddleware(limiter *IPRateLimiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		lim := limiter.GetLimiter(ip)

		if !lim.Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "Rate limit exceeded",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
