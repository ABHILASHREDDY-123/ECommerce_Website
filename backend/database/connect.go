package database

import (
	"log"
	"os"

	"github.com/ABHILASHREDDY-123/ECommerce_Website.git/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var DB *gorm.DB
func Connect() {
	err := godotenv.Load()
	if(err!=nil){
		log.Fatal("Error loading .env file")
	}
	dsn:=os.Getenv("DSN")
	db,err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if(err!=nil){
		panic("Could not connect to database")
	}else {
		log.Println("Connect successfully")
	} 
	DB  = db
	db.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Review{},
		&models.SignUps{},
		&models.Order{},
	)
}