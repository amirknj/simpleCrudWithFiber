package database

import (
	"github.com/amirknj/simpleCrud/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dns := "root:@tcp(127.0.0.1:3306)/simple_fiber"
	connection, err := gorm.Open(mysql.Open(dns), &gorm.Config{})

	if err != nil {
		panic("could not connect database")
	}
	DB = connection
	err = connection.AutoMigrate(&models.Post{})
	if err != nil {
		panic("migration fail")
	}

}
