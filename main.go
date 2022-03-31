package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {

	dsn := "root:123456@tcp(127.0.0.1:55426)/GOTEST?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Error open database")
	}
	//auto migrate table product
	//db.AutoMigrate(&Product{})

	db.Create(&Product{Code: "X24", Price: 500})
	// db.Create(&Product{Code: "E25", Price: 250})

	fmt.Println("Producto creado correctamente")

	//var product Product
	//db.First(&product, "code=?", "F24")
	//fmt.Printf("Product code: %v Price: %d\n", product.Code, product.Price)

	// update product
	// db.Model(&product).Update("Price", 340)

	// fmt.Printf("Product code: %v Price: %d\n", product.Code, product.Price)

	// id := 2
	// db.Delete(&Product{}, id)
	// fmt.Printf("Delete Product Id %d", id)
}
