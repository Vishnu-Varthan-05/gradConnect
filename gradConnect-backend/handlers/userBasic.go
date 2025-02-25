package handlers

import (
	"gradConnect-backend/database"
	"gradConnect-backend/models"
	"gradConnect-backend/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddUserInterest(c *gin.Context){
	var input struct{
		InterestIds []uint `json:"interest_ids"`
	}
	userId, _ := c.Get("userId")

	if err:= c.ShouldBindJSON(&input); err!= nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})	
		return
	}
	if len(input.InterestIds) == 0{
		c.JSON(http.StatusBadRequest, gin.H{"error":"At least one interest is required"})
		return
	}

	var interests []models.Interest
	if err:=database.DB.Where("id IN ?", input.InterestIds).Find(&interests).Error; err!= nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error":"Error fetching interests"})
		return
	}

	if len(input.InterestIds) != len(interests) {
		c.JSON(http.StatusBadRequest, gin.H{"error":"One or more interest is invalid"})
		return
	}

	var user models.User
	if err := database.DB.First(&user, userId).Error; err!= nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error":"Error retrieving user"})
		return
	}

	for _, interest := range interests{
		if !interestExists(user.Interests, interest.Id) {
			user.Interests = append(user.Interests, interest)
		}
	}

	if err := database.DB.Save(&user).Error; err!= nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save user interests"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Interests added successfully"})
}

func interestExists(interests []models.Interest, interestId uint)bool{
	for _, i := range interests{
		if interestId == i.Id{
			return true
		}
	}
	return false
}

func AddUserProfile(c *gin.Context){
	userId, _ := c.Get("userId")
	file, err := c.FormFile("photo")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No photo file provided"})	
		return
	}

	uploadDir := "./uploads"
	storage := storage.NewLocalStorage(uploadDir)

	fileUrl, err := storage.SaveFile(file)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := database.DB.First(&user, userId).Error; err!= nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error":"Error retrieving user"})
		return
	}
	user.ProfilePhotoURL = &fileUrl

	if err := database.DB.Save(&user).Error; err!= nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save user interests"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Profile photo updated successfully", "photo_url": fileUrl})
}