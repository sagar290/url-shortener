package Router

import (
	"url-shortener/Controller"

	"github.com/gin-gonic/gin"
)

func RegisterHomeRoute(r *gin.Engine) {
	// register static route
	r.Static("/assets", "./public/assets")

	r.GET("/", Controller.HomePage)
	r.POST("/PostHomePage", Controller.PostHomePage)
	r.GET("/query", Controller.QueryString)
	r.GET("/param/:name/:age", Controller.ParamString)
}
