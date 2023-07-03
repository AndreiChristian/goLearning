package main

import (
	"ginrestapi/initializers"

	"github.com/gin-gonic/gin"
)

func init() {

	initializers.LoadEnvVariables()

	initializers.ConnectToDb()

}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"hello": "world",
		})

	})

	r.Run()
}
