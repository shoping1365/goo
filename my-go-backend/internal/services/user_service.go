package services

import (
	"my-go-backend/internal/models"
)

func GetAllUsers() []models.User {
	return []models.User{
		{ID: 1, Name: "علی"},
		{ID: 2, Name: "زهرا"},
	}
}
