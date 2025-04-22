package handlers

import (
	"gin-auth-api-example/schema/response"
	"gin-auth-api-example/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Refresh(c *gin.Context) {
	rt, err := c.Cookie("refresh_token")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "no refresh token"})
		return
	}

	service, err := services.RefreshService(c, rt)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, response.RefreshResponse{
		AccessToken: service.AccessToken,
	})
}
