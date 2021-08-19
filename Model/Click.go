package Model

import (
	"time"
)

type Click struct {
	Click_id  uint `gorm:"primaryKey"`
	Url_id    int
	Referer   string `gorm:"size:255"`
	Count     int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func init() {
	// Db.AutoMigrate(Click{})
}
