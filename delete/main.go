package main

import (
	"fmt"
	"go-sql/config"
)

func main() {
	fmt.Println("Connecting to database...")
	db := config.NewDatabase()

	if db == nil {
		fmt.Println("Failed connecting to database")
		return
	}

	fmt.Println("Main function has been trigerred.")

	data := config.Student{}
	db.Unscoped().Where("id = ?", 1).Delete(&data) // Permanent Delete
	db.Where().Delete(&data)                       // Soft Delete
}
