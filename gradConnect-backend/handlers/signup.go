package handlers

import (
	"gradConnect-backend/database"
	"gradConnect-backend/models"
	"gradConnect-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {
	var input struct {
		FirstName string `json:"firstname"`
		LastName  string `json:"lastname"`
		Email     string `json:"email"`
		Password  string `json:"password"`
		Roles     []uint `json:"roles"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if len(input.Roles) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "At least one role is required"})
		return
	}

	var roles []models.Role
	if err := database.DB.Where("id IN ?", input.Roles).Find(&roles).Error; err!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error":"Error fetching roles"})
		return
	}
	

	if len(roles) != len(input.Roles){
		c.JSON(http.StatusBadRequest, gin.H{"error":"One or more roles are invalid"})
		return
	}

	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash password"})
	}
	
	user := models.User{
		FirstName: input.FirstName,
		LastName: input.LastName,
		Email: input.Email,
		Password: hashedPassword,
		Roles: roles,
	}

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
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
	c.JSON(http.StatusCreated, gin.H{"token":token})
}
