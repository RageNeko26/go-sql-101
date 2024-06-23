package config

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"fmt"
)

// Ngehasilin table 'students'
type Student struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"type:varchar(120)"`
	Age       int8
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

// # Create - Complete
// # Read / Get Data
// # Update
// # Delete
func NewDatabase() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=learning port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Something went wrong in connection database", err)
		return nil
	}

	fmt.Println("Database is connected")

	db.AutoMigrate(&Student{})
	db.Debug()

	return db
}
