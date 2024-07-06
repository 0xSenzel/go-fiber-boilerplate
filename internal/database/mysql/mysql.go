package mysql

import (
	"log"

	"github.com/0xsenzel/go-fiber-boilerplate/internal/services/user/tables"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DATABASE_URI string = "root:12345@tcp(localhost:3306)/mysql?charset=utf8mb4&parseTime=True&loc=Local"

func ConnectDB() *gorm.DB {
	var err error 

	db, err := gorm.Open(mysql.Open(DATABASE_URI), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}

	log.Println("Successfully connected to the database")
	return db
}

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&tables.User{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	log.Println("Database migration completed successfully")
	return nil
}

func CloseDB(db *gorm.DB) error {
	database, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get database connection: %v", err)
	}

	err = database.Close()
	if err != nil {
		log.Fatalf("failed to close database connection: %v", err)
	}

	log.Println("Database connection closed successfully")
	return nil
}