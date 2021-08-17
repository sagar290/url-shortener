package Services

import (
	models "url-shortener/Model"
	structs "url-shortener/Structs"
)

var User_id int64

func StoreUrl(url structs.UrlBody) (bool, string) {

	// var response models.Url

	var message string

	dbc := models.Db.Create(&models.Url{
		User_id:      int64(User.ID),
		Redirect_url: url.Redirect_url,
	})

	if dbc.Error != nil {
		message = "cant create user"
		return false, message
	}

	message = "created"

	return true, message

}

func GetUrls() []models.Url {
	var urls []models.Url

	result := models.Db.Where(&models.Url{User_id: int64(User.ID)}).Find(&urls)

	if result.Error != nil {
		return urls
	}

	return urls

}

func GetUrl(url_id int64) models.Url {
	var url models.Url

	result := models.Db.Where(&models.Url{Url_id: uint(url_id)}).First(&url)

	if result.Error != nil {
		return url
	}

	return url

}
