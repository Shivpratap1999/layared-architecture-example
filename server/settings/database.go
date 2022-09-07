package settings

import (
	"log"
	"os"
	"practice-project/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

var DBUrl = os.Getenv("GORM_DATABASE_DSN")

// this function will connect with database
func ConnectGormDB() (*gorm.DB, error) {
	var err error
	db, err = gorm.Open(mysql.Open(DBUrl), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	log.Println("Gorm DB Connected Successfully .")
	db.AutoMigrate(&models.User{})
	return db, nil
}
