package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequireRole(roles ...string) gin.HandlerFunc {
	allow := map[string]bool{}
	for _, r := range roles {
		allow[r] = true
	}
	return func(c *gin.Context) {
		role, _ := c.Get("role")
		rs, _ := role.(string)
		if !allow[rs] {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "error", "error_code": "FORBIDDEN", "message": "insufficient role", "details": []string{}})
			return
		}
		c.Next()
	}
}
