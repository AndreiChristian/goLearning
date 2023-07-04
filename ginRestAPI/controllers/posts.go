package controllers

import (
	"ginrestapi/initializers"
	"ginrestapi/models"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {

	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {

		c.JSON(400, gin.H{

			"err": result.Error,
		})

	}

	c.JSON(200, post)

}

func GetPosts(c *gin.Context) {

	var posts []models.Post

	result := initializers.DB.Find(&posts)

	if result.Error != nil {

		c.JSON(400, gin.H{

			"err": result.Error,
		})

	}

	c.JSON(200, gin.H{
		"posts": posts,
	})

}

func GetPostById(c *gin.Context) {

	var post models.Post

	index := c.Params.ByName("index")

	result := initializers.DB.First(&post, index)

	if result.Error != nil {

		c.JSON(400, gin.H{
			"err": result.Error,
		})

	}

	c.JSON(200, gin.H{
		"post": post,
	})

}
