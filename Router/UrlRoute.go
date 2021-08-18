package Router

import (
	"url-shortener/Controller"
	middleware "url-shortener/Middleware"

	"github.com/gin-gonic/gin"
)

func RegisterUrlsRoute(r *gin.Engine) {
	url := r.Group("/")

	url.Use(middleware.AuthRequired())
	{
		url.GET("/urls", Controller.GetUrls)
		url.GET("/urls/:url_id", Controller.GetUrl)
		url.PUT("/urls/:url_id", Controller.UpdateUrl)
		url.DELETE("/urls/:url_id", Controller.DeleteUrl)
		url.POST("/urls", Controller.AddUrl)
	}

	r.GET("/:url", Controller.CallbackUrl)
}
