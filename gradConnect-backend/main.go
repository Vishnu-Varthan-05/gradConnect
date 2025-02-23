package main

import (
	"gradConnect-backend/database"
	"gradConnect-backend/routes"
	"gradConnect-backend/seeders"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()
	
	seeders.Seed()

	router := gin.Default()
	router.StaticFS("/upload",http.Dir("uploads"))
	router.Use(cors.Default())
	routes.SetupRoutes(router)
	
	router.Run(":5000")
}