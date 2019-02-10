package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Product struct {
	gorm.Model
	Code string
	Price uint
}

func main() {
	DBMS     := "mysql"
	USER     := "root"
	PASS     := "password"
	PROTOCOL := "tcp(127.0.0.1:3306)"
	DBNAME   := "gorm"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "L1212", Price: 1000})

	// Read
	var product Product
	db.First(&product, 1)
	fmt.Printf("%+v\n", product)

	db.First(&product, "code = ?", "L1212")
	fmt.Printf("%+v\n", product)

	// Update
	db.Model(&product).Update("Price", 2000)
	fmt.Printf("%+v\n", product)

	// Delete
	db.Delete(&product)
}