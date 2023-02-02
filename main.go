package main

import (
	"apiRest/controllers"
	"apiRest/database"
	"apiRest/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	database.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/posts", controllers.PostCreate)
	r.GET("/posts", controllers.GetPosts)
	r.GET("/posts/:id", controllers.GetPostById)
	r.PUT("/posts/:id", controllers.UpdatePost)
	r.DELETE("/posts/:id", controllers.DeletePost)

	r.Run()
}
