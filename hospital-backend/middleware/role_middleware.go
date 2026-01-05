package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func RoleMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		rawRoles, exists := c.Get("roles")
		if !exists {
			c.JSON(http.StatusForbidden, gin.H{"error": "Acceso denegado"})
			c.Abort()
			return
		}

		userRoles, ok := rawRoles.([]string)
		if !ok {
			c.JSON(http.StatusForbidden, gin.H{"error": "Acceso denegado"})
			c.Abort()
			return
		}

		for _, userRole := range userRoles {
			for _, allowed := range allowedRoles {
				if strings.EqualFold(userRole, allowed) {
					c.Next()
					return
				}
			}
		}

		c.JSON(http.StatusForbidden, gin.H{"error": "No tienes permisos"})
		c.Abort()
	}
}
