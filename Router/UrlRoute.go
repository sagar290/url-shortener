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
		url.POST("/urls", Controller.AddUrl)
	}

	r.GET("/:url", Controller.CallbackUrl)
}
