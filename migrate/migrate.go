package main

import (
	"apiRest/database"
	"apiRest/initializers"
	"apiRest/models"
)

func init() {
	initializers.LoadEnv()
	database.ConnectToDB()
}

func main() {
	database.DB.AutoMigrate(&models.Post{})
}
