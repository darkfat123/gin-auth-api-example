package services

import (
	"errors"
	"gin-auth-api-example/redis"
	"gin-auth-api-example/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type refreshResult struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func RefreshService(c *gin.Context, rt string) (*refreshResult, error) {
	token, err := utils.VerifyRefreshToken(rt)
	if err != nil || !token.Valid {
		return nil, errors.New("invalid refresh token")
	}

	userID := int(token.Claims.(jwt.MapClaims)["user_id"].(float64))

	val, err := redis.GetData(strconv.Itoa(userID))
	if err != nil || val != rt {
		return nil, errors.New("refresh token mismatch or not found")
	}

	err = redis.DeleteData(strconv.Itoa(userID))
	if err != nil {
		return nil, errors.New("failed to invalidate old refresh token")
	}

	newAccessToken, err := utils.GenerateAccessToken(userID)
	if err != nil {
		return nil, errors.New("failed to generate access token")
	}

	newRefreshToken, err := utils.GenerateRefreshToken(userID)
	if err != nil {
		return nil, errors.New("failed to generate new refresh token")
	}

	expirationTime := 3600 * 24
	err = redis.SetData(strconv.Itoa(userID), newRefreshToken, time.Duration(expirationTime)*time.Second)
	if err != nil {
		return nil, errors.New("failed to store new refresh token")
	}

	return &refreshResult{
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
	}, nil
}
