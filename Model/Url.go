package Model

import (
	"time"
)

type Url struct {
	Url_id       uint `gorm:"primaryKey"`
	User_id      int64
	Redirect_url string `gorm:"size:255"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func init() {
	// Db.AutoMigrate(Url{})
}
