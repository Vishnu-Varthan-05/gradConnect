package seeders

import (
	"gradConnect-backend/database"
	"gradConnect-backend/models"
)

func Seed() {
	roles := []models.Role{
		{Name: "Student"},
		{Name: "Alumni"},
		{Name: "Admin"},
	}
	for _, role := range roles{
		database.DB.FirstOrCreate(&role, models.Role{Name: role.Name})
	}
}