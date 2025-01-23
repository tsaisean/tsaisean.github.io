package main

import (
	"fmt"
	"iHR/config"
	"iHR/db"
	"iHR/db/models"
	"log"
	"reflect"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	db.Connect(&cfg.Database)
	// Add column
	// addColumn(&models.Employee{}, "Salary")

	// Alter column
	alterColumn(&models.Employee{}, "salary")

	// Rename column
	//renameColumn(&models.Employee{}, "first_name2", "first_name")

	// Drop column
	//dropColumn(&models.Employee{}, "first_name")
	// Change Column Type

	// Drop Table
	// dropTable(&models.Employee{})
}

func addColumn(table interface{}, structFieldName string) {
	if err := db.DB.Migrator().AddColumn(table, structFieldName); err != nil {
		fmt.Printf("Failed to add column: %v\n", err)
		return
	}
	fmt.Println("Successfully added column: ", structFieldName)
}

func alterColumn(table interface{}, structFieldName string) {
	if err := db.DB.Migrator().AlterColumn(table, structFieldName); err != nil {
		fmt.Printf("Failed to alter column: %v\n", err)
		return
	}
	fmt.Println("Successfully alter column: ", structFieldName)
}

func renameColumn(table interface{}, oldName, newName string) {
	if err := db.DB.Migrator().RenameColumn(table, oldName, newName); err != nil {
		fmt.Printf("Failed to rename column: %v\n", err)
		return
	}
	fmt.Println("Successfully renamed column from", oldName, "to", newName)
}

func dropColumn(table interface{}, colName string) {
	if err := db.DB.Migrator().DropColumn(table, colName); err != nil {
		fmt.Printf("Failed to dropped column: %v\n", err)
		return
	}
	fmt.Println("Successfully dropped column:", colName)
}

func dropTable(table interface{}) {
	if err := db.DB.Migrator().DropTable(table); err != nil {
		fmt.Printf("Failed to dropped table: %v\n", err)
		return
	}
	fmt.Println("Successfully dropped table:", reflect.TypeOf(table).Elem().Name())
}
