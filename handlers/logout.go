package handlers

import (
	"gin-auth-api-example/redis"
	"gin-auth-api-example/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	rt, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to get refresh token"})
		return
	}

	claims, err := utils.ParseToken(rt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse refresh token"})
		return
	}

	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user_id in token claims"})
		return
	}

	userID := strconv.Itoa(int(userIDFloat))
	redis.DeleteData(userID)
	c.SetCookie("refresh_token", "", -1, "/", "", true, true)

	c.JSON(http.StatusOK, gin.H{"message": "logged out"})
}
