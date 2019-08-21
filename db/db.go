package db

import (
	"fmt"
	"newfeed/models"

	"github.com/jinzhu/gorm"
	//mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

//InitDatabase blala
func InitDatabase() *gorm.DB {
	//open a db connection
	var err error
	db, err = gorm.Open("mysql", "root:1@/newfeeddb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	//Migrate the schema
	db.AutoMigrate(&models.NewfeedModel{})
	return db
}
