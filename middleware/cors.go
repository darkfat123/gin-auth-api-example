package middleware

import "github.com/gin-gonic/gin"

// ปรับให้ไม่มีการตอบกลับ OPTIONS
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // หรือระบุ domain เช่น https://example.com
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		c.Next()
	}
}
