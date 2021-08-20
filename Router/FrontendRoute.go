package Router

import (
	"url-shortener/Controller"

	"github.com/gin-gonic/gin"
)

func RegisterFrontendRoute(r *gin.Engine) {
	r.GET("/", Controller.HomePage)
	r.GET("/url/:slug", Controller.SingleUrlpage)
	r.GET("/login", Controller.LoginPage)
	r.GET("/signup", Controller.SignupPage)
}
