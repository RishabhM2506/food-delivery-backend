package middleware

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	rds "github.com/redis/go-redis/v9"
)

func SlidingWindowRateLimit(rc *rds.Client, limit int, window time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := fmt.Sprintf("ratelimit:ip:%s", c.ClientIP())
		cnt, _ := rc.Incr(context.Background(), key).Result()
		if cnt == 1 {
			_ = rc.Expire(context.Background(), key, window).Err()
		}
		if int(cnt) > limit {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"status": "error", "error_code": "RATE_LIMITED", "message": "too many requests", "details": []string{}})
			return
		}
		c.Next()
	}
}
