package providers

import (
	"github.com/imanimen/foodrate/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type IDatabase interface {

}

type Database struct {
	Connection *gorm.DB
	Config     IConfig
}

func NewDatabase(config IConfig) IDatabase {
	dsn := config.Get("dsn")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error connecting db")
	}

	db.AutoMigrate(&models.Place{})
	db.AutoMigrate(&models.User{})

	return &Database{
		Connection: db,
		Config:     config,
	}
}
