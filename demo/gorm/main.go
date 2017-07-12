package main

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Product struct {
	gorm.Model

	Code string
	Price uint
}

func main() {
	db, err := gorm.Open("mysql", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&Product{})

	db.Create(&Product{Code: "foo", Price: 1000})

	var product Product
	dberr := db.First(&product, 6)
	if dberr != nil {
		log.Println(product.Error)
	} else {
		log.Println("id:", product.ID)
	}

	db.First(&product, "code = ?", "foo")
	log.Println("id:", product.ID)
	log.Println("code:", product.Code)
	log.Println("price:", product.Price)

	db.Model(&product).Update("Price", 2000)
	db.First(&product, "code = ?", "foo")

	log.Println("price:", product.Price)

	db.Delete(&product)
}
