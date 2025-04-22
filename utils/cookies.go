package utils

import "github.com/gin-gonic/gin"

func SetRefreshCookie(c *gin.Context, token string) {
	c.SetCookie("refresh_token", token, 24*3600, "/", "", true, true)
}

func ClearRefreshCookie(c *gin.Context) {
	c.SetCookie("refresh_token", "", -1, "/", "", true, true)
}
