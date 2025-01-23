package main

import (
	"github.com/gin-gonic/gin"
	"iHR/config"
	"iHR/db"
	. "iHR/db/models"
	"iHR/route"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	db.Connect(&cfg.Database)

	// AutoMigrations
	err = db.DB.AutoMigrate(&Employee{})
	if err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	}

	r := gin.Default()

	route.RegisterEmployeeRoutes(r)

	r.Run()
}
