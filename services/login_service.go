package services

import (
	"errors"
	"gin-auth-api-example/database"
	"gin-auth-api-example/model"
	"gin-auth-api-example/redis"
	"gin-auth-api-example/schema/request"
	"gin-auth-api-example/utils"
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type loginResult struct {
	AccessToken  string
	RefreshToken string
}

func LoginService(c *gin.Context, req *request.LoginRequest) (*loginResult, error) {
	query := `SELECT id, password FROM users WHERE username = $1`
	var user model.Users
	err := database.DB.GetContext(c, &user, query, req.Username)
	if err != nil {
		log.Println("database err:", err)
		return nil, errors.New("invalid username or password")
	}

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		return nil, errors.New("invalid username or password")
	}

	accessToken, err := utils.GenerateAccessToken(user.ID)
	if err != nil {
		return nil, err
	}

	refreshToken, err := utils.GenerateRefreshToken(user.ID)
	if err != nil {
		return nil, err
	}

	key := strconv.Itoa(user.ID)
	err = redis.SetData(key, refreshToken, 24*time.Hour)

	if err != nil {
		log.Println("Redis err:", err)
	}

	return &loginResult{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
