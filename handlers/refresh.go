package handlers

import (
	"gin-auth-api-example/schema/response"
	"gin-auth-api-example/services"
	"net/http"
	"os"
	"strconv"

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

	c.JSON(http.StatusOK, response.RefreshResponse{
		AccessToken: service.AccessToken,
	})
}
