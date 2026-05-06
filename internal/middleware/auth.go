package middleware

import (
	"context"
	"net/http"
	"strings"

	"food-delivery-backend/pkg/config"

	"github.com/gin-gonic/gin"
	rds "github.com/redis/go-redis/v9"
)

func JWTAuthMiddleware(cfg *config.Config, rc *rds.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if !strings.HasPrefix(auth, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "error", "error_code": "UNAUTHORIZED", "message": "invalid bearer", "details": []string{}})
			return
		}
		tok := strings.TrimPrefix(auth, "Bearer ")
		if tok == "" || cfg.JWT.Secret == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "error", "error_code": "UNAUTHORIZED", "message": "invalid token", "details": []string{}})
			return
		}
		sid := c.GetHeader("X-Session-ID")
		if sid == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "error", "error_code": "UNAUTHORIZED", "message": "missing session", "details": []string{}})
			return
		}
		if _, err := rc.Get(context.Background(), "session:"+sid).Result(); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "error", "error_code": "UNAUTHORIZED", "message": "session expired", "details": []string{}})
			return
		}
		c.Set("sid", sid)
		c.Next()
	}
}
