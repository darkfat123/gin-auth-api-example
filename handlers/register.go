package handlers

import (
	"gin-auth-api-example/schema/request"
	"gin-auth-api-example/schema/response"
	"gin-auth-api-example/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var req request.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	id, err := services.RegisterService(c, &req)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, response.RegisterResponse{
		Message: "Registed Successfully!",
		ID:      id,
	})
}
