package Controller

import (
	"fmt"
	"net/http"
	models "url-shortener/model"

	"github.com/gin-gonic/gin"
)

type Formbody struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}

func HomePage(c *gin.Context) {

	// fetch database
	var user models.User
	db := models.Db
	db.Model(&models.User{}).First(&user)

	fmt.Println(user.Name)

	c.HTML(http.StatusOK, "pages/home", gin.H{
		"title": "Page file title!!",
		"add": func(a int, b int) int {
			return a + b
		},
	})
}

func PostHomePage(c *gin.Context) {
	var formbody Formbody

	c.BindJSON(&formbody)

	c.JSON(200, gin.H{
		"name": formbody.Name,
		"age":  formbody.Age,
	})
}

func QueryString(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")
	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}

func ParamString(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")
	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}
