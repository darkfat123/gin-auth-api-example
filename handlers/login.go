package handlers

import (
	"gin-auth-api-example/schema/request"
	"gin-auth-api-example/schema/response"
	"gin-auth-api-example/services"
	"net/http"
	"os"
	"strconv"

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

	maxAge, err := strconv.Atoi(os.Getenv("REFRESH_TOKEN_MAX_AGE"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid REFRESH_TOKEN_MAX_AGE value"})
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "refresh_token",
		Value:    service.RefreshToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   maxAge,
	})

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "csrf_token",
		Value:    service.CsrfToken,
		Path:     "/",
		HttpOnly: false,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   maxAge,
	})

	c.JSON(http.StatusOK, response.LoginResponse{
		AccessToken: service.AccessToken,
	})
}
