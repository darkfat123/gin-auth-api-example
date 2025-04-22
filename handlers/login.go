package handlers

import (
	"gin-auth-api-example/schema/request"
	"gin-auth-api-example/schema/response"
	"gin-auth-api-example/services"
	"gin-auth-api-example/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var req request.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	service, err := services.LoginService(c, &req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	utils.SetRefreshCookie(c, service.RefreshToken)

	c.JSON(http.StatusOK, response.LoginResponse{
		AccessToken: service.AccessToken,
	})
}
