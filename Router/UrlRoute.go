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
		url.GET("api/urls", Controller.GetUrls)
		url.GET("api/urls/:url_id", Controller.GetUrl)
		url.PUT("api/urls/:url_id", Controller.UpdateUrl)
		url.DELETE("api/urls/:url_id", Controller.DeleteUrl)
		url.POST("api/urls", Controller.AddUrl)
	}

	r.GET("/:slug", Controller.CallbackUrl)
}
