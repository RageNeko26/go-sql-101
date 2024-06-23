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

	// Updates()
	db.Where("id IN ?", []uint{3, 4, 5, 6, 7, 8}).Updates(&config.Student{
		Name: "DONI",
	})

}
