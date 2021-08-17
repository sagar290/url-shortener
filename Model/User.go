package Model

import (
	"time"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:255"`
	Email     string `gorm:"size:255"`
	Password  string `gorm:"size:255"`
	Status    int16  `gorm:"default:0"`
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func init() {
	// Db.AutoMigrate(User{})
}
