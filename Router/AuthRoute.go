package Router

import (
	"url-shortener/Controller"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoute(r *gin.Engine) {
	// register static route
	r.Static("/assets", "./public/assets")

	// r.GET("/", Controller.HomePage)
	r.POST("api/login", Controller.Login)
	r.POST("api/register", Controller.Register)
}
