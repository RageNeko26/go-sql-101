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

	// Create new data
	student_1 := []config.Student{
		{
			Name: "Joko",
			Age:  8,
		},
		{
			Name: "Sheryl",
			Age:  11,
		},
		{
			Name: "Anton",
			Age:  20,
		},
	}

	db.Create(&student_1)

	fmt.Println("Data has been created")

}
