package Model

import (
	"time"
)

type Url struct {
	Url_id       uint      `gorm:"primaryKey" json:"url_id"`
	User_id      int64     `json:"user_id"`
	Slug         string    `json:"slug"`
	TotalCount   int       `json:"total_count"`
	Redirect_url string    `gorm:"size:255" json:"redirect_url"`
	Clicks       []Click   `gorm:"foreignKey:Url_id;references:Url_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func init() {
	// Db.AutoMigrate(Url{})
}
