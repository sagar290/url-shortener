package Services

import (
	models "url-shortener/Model"
	structs "url-shortener/Structs"
)

var Token string

var User models.User

func CheckLoginAttempt(body structs.LoginBody) (bool, string, models.User) {

	hash := GenerateHash(body.Password)

	var user models.User
	var message string

	models.Db.Where(&models.User{Email: body.Email, Password: hash}).First(&user)

	if user.Email != body.Email {
		message = "user not found"
		return false, message, user
	}

	message = "user found"

	return true, message, user
}

func StoreUser(user structs.RegisterBody) (bool, string, models.User) {

	var response models.User

	var exist_user models.User
	var message string

	models.Db.Where(&models.User{Email: user.Email}).First(&exist_user)

	if exist_user.Email == user.Email {
		message = "user already exist"
		return false, message, response
	}

	hash := GenerateHash(user.Password)

	dbc := models.Db.Create(&models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: hash,
		Token:    GenerateToken(100),
	})

	if dbc.Error != nil {
		message = "cant create user"
		return false, message, response
	}

	models.Db.Where(&models.User{Email: user.Email}).First(&response)

	message = "created"

	return true, message, response
}

func SetToken(token string) {
	Token = token
}

func VerifyToken(token string) bool {

	result := models.Db.Where(&models.User{Token: token}).First(&User)

	if result.Error != nil {
		return false
	}

	if User.Token == "" {
		return false
	}

	// asign the user id
	// User_id = int64(User.ID)

	return true
}
