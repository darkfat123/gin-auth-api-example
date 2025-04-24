package handlers

import (
	"gin-auth-api-example/schema/request"
	"gin-auth-api-example/schema/response"
	"gin-auth-api-example/services"
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

	c.SetCookie("refresh_token", service.RefreshToken, 3600*24, "/", "", true, true)
	c.SetCookie("csrf_token", service.CsrfToken, 3600*24, "/", "", true, false)

	c.JSON(http.StatusOK, response.LoginResponse{
		AccessToken: service.AccessToken,
	})
}
