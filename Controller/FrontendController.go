package Controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {

	c.HTML(http.StatusOK, "pages/home", gin.H{
		"title": "home",
	})
}

func SingleUrlpage(c *gin.Context) {
	c.HTML(http.StatusOK, "pages/singleurl", gin.H{
		"title": "singleurl",
	})
}

func LoginPage(c *gin.Context) {

	c.HTML(http.StatusOK, "pages/login", gin.H{
		"title": "login",
	})
}

func SignupPage(c *gin.Context) {

	c.HTML(http.StatusOK, "pages/signup", gin.H{
		"title": "signup",
	})
}
