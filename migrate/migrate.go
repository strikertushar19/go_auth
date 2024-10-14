package main

import (
	"go_auth/initializers"
	"go_auth/models"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()

}

func main() {
	
     initializers.DB.AutoMigrate(&models.User{})
}