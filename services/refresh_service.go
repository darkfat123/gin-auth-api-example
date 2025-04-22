package services

import (
	"errors"
	"gin-auth-api-example/redis"
	"gin-auth-api-example/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type refreshResult struct {
	AccessToken string `json:"access_token"`
}

func RefreshService(c *gin.Context, rt string) (*refreshResult, error) {
	token, err := utils.VerifyRefreshToken(rt)
	if err != nil || !token.Valid {
		return nil, errors.New("invalid refresh token")
	}

	userID := int(token.Claims.(jwt.MapClaims)["user_id"].(float64))

	val, err := redis.GetData(strconv.Itoa(userID))
	if err != nil || val != rt {
		return nil, errors.New("token not found")
	}

	newAccessToken, _ := utils.GenerateAccessToken(userID)

	return &refreshResult{
		AccessToken: newAccessToken,
	}, nil
}
