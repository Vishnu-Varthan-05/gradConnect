package handlers

import (
	"gradConnect-backend/database"
	"gradConnect-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllInterests(c *gin.Context) {
	var interests []models.Interest
	if err := database.DB.Find(&interests).Error; err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error":"Error fetching interests"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"interests":interests})
}

func AddInterest(c *gin.Context) {
	var input struct{
		Name string
	}
	if err := c.ShouldBindJSON(&input); err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var interest models.Interest = models.Interest{
		Name: input.Name,
	}
	if err := database.DB.Create(&interest).Error; err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error":"Unable to add interest"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"interest": interest})
}