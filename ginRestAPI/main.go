package main

import (
	"ginrestapi/controllers"
	"ginrestapi/initializers"

	"github.com/gin-gonic/gin"
)

func init() {

	initializers.LoadEnvVariables()

	initializers.ConnectToDb()

}

func main() {
	r := gin.Default()

	r.POST("/", controllers.CreatePost)

	r.GET("/", controllers.GetPosts)

	r.GET("posts/:id", controllers.GetPostById)
}
