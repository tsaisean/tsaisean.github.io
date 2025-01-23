package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"iHR/config"
	"log"
	"sync"
)

var (
	DB   *gorm.DB
	once sync.Once
)

func Connect(cfg *config.Database) {
	if DB != nil {
		fmt.Println("MySQL is already connected!")
		return
	}

	once.Do(func() {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%s&loc=%s", cfg.Username, cfg.Password,
			cfg.Host, cfg.Port, cfg.DBName, "utf8mb4", "True", "Local")
		var err error
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("Failed to connect to MySQL: %v", err)
		}

		fmt.Println("Connected to MySQL successfully!")
	})
}
