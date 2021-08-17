package Controller

import (
	structs "url-shortener/Structs"

	services "url-shortener/Services"

	models "url-shortener/Model"

	"github.com/gin-gonic/gin"
)

func AddUrl(c *gin.Context) {
	var urlBody structs.UrlBody

	if err := c.ShouldBindJSON(&urlBody); err != nil {
		c.JSON(422, gin.H{"error": err.Error()})
		return
	}

	results, message := services.StoreUrl(urlBody)

	if results == false {
		c.JSON(400, gin.H{
			"message": message,
			"data":    []gin.H{},
		})
	}

	c.JSON(200, gin.H{
		"message": message,
		"data":    []gin.H{},
	})
}

func GetUrls(c *gin.Context) {

	var urls []models.Url

	urls = services.GetUrls()

	c.JSON(200, gin.H{
		"message": "",
		"data":    urls,
	})
}

func CallbackUrl(c *gin.Context) {
	url := c.Param("url")

	c.JSON(200, gin.H{
		"message": url,
		"data":    []gin.H{},
	})
}
