package models

import "time"

/* Note: GORM's AutoMigrate will not work for rename, drop a column or incompatible column type changes. Please do manual migration when necessary.*/

type Employee struct {
	ID           uint    `gorm:"primary_key"`
	FirstName    string  `gorm:"size:50"`
	LastName     string  `gorm:"size:50"`
	Email        string  `gorm:"size:100"`
	Position     string  `gorm:"size:50"`
	SupervisorID *uint64 `gorm:"index"`
	Salary       uint
	CreatedAt    time.Time // Auto handled by GORM
}
