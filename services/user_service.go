package services

import (
	"errors"
	"gin-auth-api-example/database"
	"gin-auth-api-example/model"

	"github.com/gin-gonic/gin"
)

func GetUserByIDService(c *gin.Context, id string) (*model.Users, error) {
	var user model.Users

	query := `SELECT * FROM users where id = $1`
	err := database.DB.GetContext(c, &user, query, id)
	if err != nil {
		return nil, errors.New("data not found")
	}

	return &user, nil
}
