package middleware

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func SecurityHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("X-Frame-Options", "DENY")
		c.Writer.Header().Set("X-Content-Type-Options", "nosniff")
		c.Writer.Header().Set("X-XSS-Protection", "1; mode=block")
		c.Writer.Header().Set("Referrer-Policy", "no-referrer")
		c.Writer.Header().Set("Strict-Transport-Security", "max-age=63072000; includeSubDomains; preload")
		c.Next()
	}
}

func CSRFMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		csrfFromCookie, err := c.Cookie("csrf_token")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "missing CSRF cookie"})
			return
		}

		csrfFromHeader := c.GetHeader("X-CSRF-Token")
		if csrfFromHeader == "" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "missing CSRF header"})
			return
		}

		decodedHeader, err := url.QueryUnescape(csrfFromHeader)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "invalid csrf header encoding"})
			return
		}

		if decodedHeader != csrfFromCookie {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "CSRF token mismatch"})
			return
		}

		c.Next()
	}
}
