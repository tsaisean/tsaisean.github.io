package main

import (
	"encoding/csv"
	"fmt"
	"iHR/config"
	"iHR/db"
	"iHR/db/models"
	"log"
	"os"
	"strings"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	db.Connect(&cfg.Database)

	file, err := os.Open("db/models/employee_seed.csv")
	if err != nil {
		log.Fatalf("Failed to open CSV file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Failed to parse CSV: %v", err)
	}

	for i, record := range records {
		if i == 0 {
			// Skip header row
			continue
		}

		employee := models.Employee{
			FirstName: strings.TrimSpace(record[0]),
			LastName:  strings.TrimSpace(record[1]),
			Email:     strings.TrimSpace(record[2]),
			Position:  strings.TrimSpace(record[3]),
		}

		if err := db.DB.Create(&employee).Error; err != nil {
			log.Fatalf("Failed to create employee: %v", err)
		} else {
			log.Printf("Seed employee: %v", employee)
		}
	}

	fmt.Println("Seeding successfully!")
}
