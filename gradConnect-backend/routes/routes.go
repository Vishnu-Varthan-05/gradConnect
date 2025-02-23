package routes

import (
	"fmt"
	"gradConnect-backend/handlers"
	"gradConnect-backend/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine){
	router.POST("/signup", handlers.Signup)
	router.POST("/login", handlers.Login)

	router.GET("/interest", handlers.GetAllInterests)
	router.POST("/interest", handlers.AddInterest)

	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	protected.GET("/hello", func(ctx *gin.Context) {
		fullname, _ := ctx.Get("fullname")  
    	roles, _ := ctx.Get("roles")
		message := fmt.Sprintf("Hello mr %v with role %v", fullname, roles)
		ctx.JSON(http.StatusOK, gin.H{"message": message})
	})
	
}