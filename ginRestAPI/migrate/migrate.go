package main

import (
	"fmt"
	"ginrestapi/initializers"
	"ginrestapi/models"
)

func init() {

	initializers.LoadEnvVariables()

	initializers.ConnectToDb()

}

func main() {

	initializers.DB.AutoMigrate(&models.Post{})

	fmt.Println("Automigrated data succesfully")

}
