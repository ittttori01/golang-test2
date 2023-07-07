package initializers

import "example.com/jwt/models"

func SyncDatabase(){
	DB.AutoMigrate(&models.User{})
}