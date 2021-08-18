package Structs

var User_id string

type UrlBody struct {
	Redirect_url string `json:"redirect_url" binding:"required"`
	Slug         string `json:"slug" binding:"required"`
}

type UrlResponse struct {
	Url_id       int64  `json:"url_id"`
	User_id      int64  `json:"user_id"`
	Redirect_url string `json:"redirect_url"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}
