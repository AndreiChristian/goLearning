package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	// db.Create(&Product{Code: "D42", Price: 100})

	var product Product
	// db.

	// fmt.Println(product.ID, product.Code, product.Price)

	// db.Model(&product).Updates(Product{Price: 300, Code: "F42"})

	// db.Model(&product).Update("Price", 200)

	// db.Delete(&product, 1)

	db.First(&product, 2)

	fmt.Println(product.ID, product.Code, product.Price)
}
