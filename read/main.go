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

	student := []config.Student{}

	// ["a", "b", "c"]
	// [1, 5, 10]
	// [false, true, true]
	// [config.Student{}, config.Student{},]

	// Mysql: Support LIKE
	// Postgres: Support ILIKE

	// Syntax Where(nama_kolom, nilai_yang_dicari)
	// searchQuery := "%" + "jo" + "%"

	db.Find(&student)

	for _, data := range student {
		fmt.Println("ID:", data.ID)
	}

}
