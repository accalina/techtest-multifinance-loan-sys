package infra

import (
	"log"
	"mf-loan/entity"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	// Migrate the schema
	db.AutoMigrate(&entity.DetailCustomer{}, &entity.Tenor{}, &entity.TransactionDetail{})

	return db
}
