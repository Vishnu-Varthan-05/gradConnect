package handlers

import (
	"gradConnect-backend/database"
	"gradConnect-backend/models"
	"gradConnect-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var input struct{
		Email string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&input); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}
	var user models.User
	if err := database.DB.Where("email = ?", input.Email).Preload("Roles").First(&user).Error; err!= nil{
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	if err := utils.CheckPassword(user.Password, input.Password); err!= nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	var rolesName []string;
	for _, role := range user.Roles{
		rolesName = append(rolesName, role.Name)
	}

	token, err := utils.GenerateJWT(user.ID, user.FirstName + " " + user.LastName, rolesName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token":token})
}