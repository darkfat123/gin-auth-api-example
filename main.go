package main

import (
	"gin-auth-api-example/database"
	"gin-auth-api-example/handlers"
	"gin-auth-api-example/middleware"
	"gin-auth-api-example/redis"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading .env file: %v\n", err)
	}

	database.InitDB()
	defer database.DB.Close()

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	redis.InitRedis()

	router := gin.Default()

	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.SecurityHeaders())

	auth := router.Group("/auth")
	{
		auth.POST("/register", handlers.Register)
		auth.POST("/login", handlers.Login)
		auth.POST("/refresh", middleware.CSRFMiddleware(), handlers.Refresh)
		auth.POST("/logout", handlers.Logout)
	}
	api := router.Group("/api", middleware.JWTAuthMiddleware())
	{
		api.GET("/user/:id", handlers.GetUserByID)
	}

	router.Run()
}
