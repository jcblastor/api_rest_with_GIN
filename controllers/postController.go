package controllers

import (
	"apiRest/database"
	"apiRest/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {
	// get data of req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	post := models.Post{Title: body.Title, Body: body.Body}

	result := database.DB.Create(&post)
	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(201, gin.H{
		"post": post,
	})
}

func GetPosts(c *gin.Context) {
	posts := []models.Post{}

	database.DB.Find(&posts)

	c.JSON(200, gin.H{
		"post": posts,
	})
}

func GetPostById(c *gin.Context) {
	id := c.Param("id")

	post := models.Post{}

	database.DB.First(&post, id)

	c.JSON(200, gin.H{
		"post": post,
	})
}

func UpdatePost(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	// find de post
	post := models.Post{}
	database.DB.First(&post, id)

	// update post
	database.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	// response
	c.JSON(200, gin.H{
		"post": post,
	})
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")

	//delete post
	database.DB.Delete(&models.Post{}, id)

	// response
	c.JSON(200, gin.H{
		"msg": fmt.Sprintf("Delete post with id: %s", id),
	})
}
