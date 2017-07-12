package main

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Product struct {
	gorm.Model
	Code    string
	Price   uint
	StoreID uint
}

type Store struct {
	gorm.Model
	Products []Product
	Name     string
}

func main() {
	db, err := gorm.Open("mysql", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&Store{})
	db.AutoMigrate(&Product{})

	db.Create(&Store{Name: "Random Store"})
	db.Create(&Product{Code: "bar", Price: 2000, StoreID: 1})
	db.Create(&Product{Code: "foo", Price: 1000, StoreID: 1})

	var store Store
	products := make([]Product, 0)
	db.First(&store, "name = ?", "Random Store")
	log.Println("store:", store.Name)

	db.Model(&store).Related(&products)
	for _, product := range products {
		log.Println(product.Code, product.Price)
	}

	/*
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
	*/
}
