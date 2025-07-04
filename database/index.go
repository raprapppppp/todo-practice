package db

import (
	"todo-practice/models/todos"
	"todo-practice/models/users"

	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB is the GORM database client instance
var Database *gorm.DB

func ConnectionDB() error {
	var err error
	var dsn = "host=localhost user=postgres password=1234 dbname=simple-todos port=5432 sslmode=disable"

	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	Database.AutoMigrate(&users.Users{})
	Database.AutoMigrate(&todos.Todos{})

	log.Println("Database connected successfully!")
	log.Println("Database migration complete.")
	return nil

}
