package helpers

import (
	"fmt"
	"kautsarhasby/ewallet-ums/internal/models"
	"log"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupDatabase() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		GetEnv("DB_USER", ""),
		GetEnv("DB_PASSWORD", ""),
		GetEnv("DB_HOST", "127.0.0.1"),
		GetEnv("DB_PORT", "3306"),
		GetEnv("DB_NAME", ""),
	)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connect to Database")
	}

	logrus.Info("Successfully Connect to Database")

	DB.AutoMigrate(&models.User{}, &models.UserSession{})
}
