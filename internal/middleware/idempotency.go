package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	rds "github.com/redis/go-redis/v9"
)

func IdempotencyMiddleware(rc *rds.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.GetHeader("Idempotency-Key")
		if key == "" {
			c.Next()
			return
		}
		rkey := "idempotency:" + key
		if v, _ := rc.Get(context.Background(), rkey).Result(); v != "" {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": "success", "data": v})
			return
		}
		c.Next()
		_ = rc.Set(context.Background(), rkey, "ok", 24*time.Hour).Err()
	}
}
