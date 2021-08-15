package Model

import (
	"time"
	services "url-shortener/Services"
	structs "url-shortener/Structs"
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

func CheckLoginAttempt(body structs.LoginBody) (bool, string, User) {

	hash := services.GenerateHash(body.Password)

	var user User
	var message string

	Db.Where(&User{Email: body.Email, Password: hash}).First(&user)

	if user.Email != body.Email {
		message = "user not found"
		return false, message, user
	}

	message = "user found"

	return true, message, user
}

func StoreUser(user structs.RegisterBody) (bool, string, User) {

	var response User

	var exist_user User
	var message string

	Db.Where(&User{Email: user.Email}).First(&exist_user)

	if exist_user.Email == user.Email {
		message = "user already exist"
		return false, message, response
	}

	hash := services.GenerateHash(user.Password)

	Db.Create(&User{
		Name:     user.Name,
		Email:    user.Email,
		Password: hash,
		Token:    services.GenerateToken(100),
	})

	Db.Where(&User{Email: user.Email}).First(&response)

	message = "created"

	return true, message, response
}
