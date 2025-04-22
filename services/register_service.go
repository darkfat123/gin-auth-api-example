package services

import (
	"database/sql"
	"errors"
	"gin-auth-api-example/database"
	"gin-auth-api-example/model"
	"gin-auth-api-example/schema/request"
	"gin-auth-api-example/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func RegisterService(c *gin.Context, req *request.RegisterRequest) (int, error) {
	query := `SELECT * FROM users WHERE email = $1 OR username = $2`
	var user model.Users
	err := database.DB.GetContext(c, &user, query, req.Email, req.Username)

	if err == nil {
		return 0, errors.New("email or username already exists")
	}

	if !errors.Is(err, sql.ErrNoRows) {
		return 0, err
	}

	hashPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return 0, err
	}

	insertQuery := `
	INSERT INTO users (email, username, password, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5) 
	RETURNING id
	`

	now := time.Now()
	var id int
	err = database.DB.GetContext(c, &id, insertQuery, req.Email, req.Username, hashPassword, now, now)
	if err != nil {
		return 0, err
	}

	return id, nil

}
