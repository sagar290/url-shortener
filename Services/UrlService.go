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
		Slug:         url.Slug,
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

func GetUrl(url_id int) models.Url {
	var url models.Url

	result := models.Db.Where(&models.Url{Url_id: uint(url_id)}).First(&url)

	if result.Error != nil {
		return url
	}

	return url

}

func GetUrlBySlug(slug string) models.Url {
	var url models.Url

	result := models.Db.Where(&models.Url{Slug: slug}).First(&url)

	if result.Error != nil {
		return url
	}

	return url

}

func UpdateUrl(url_id int, body structs.UrlBody) (bool, string) {
	// var url models.Url
	var message string

	if body.Redirect_url != "" {

		result := models.Db.Where(&models.Url{Url_id: uint(url_id)}).Update("redirect_url", body.Redirect_url)
		if result.Error != nil {
			message = "cant update"
			return false, message
		}
	}

	if body.Slug != "" {

		result := models.Db.Where(&models.Url{Url_id: uint(url_id)}).Update("slug", body.Slug)
		if result.Error != nil {
			message = "cant update"
			return false, message
		}
	}

	message = "updated"

	return true, message

}

func DeleteUrl(url_id int) (bool, string) {
	var url models.Url
	var message string

	result := models.Db.Where(&models.Url{Url_id: uint(url_id)}).Delete(&url)
	if result.Error != nil {
		message = "cant delete"
		return false, message
	}

	message = "deleted"

	return true, message

}
