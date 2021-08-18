package Controller

import (
	"fmt"
	"strconv"
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

func GetUrl(c *gin.Context) {
	url_id := c.Param("url_id")
	var url models.Url

	// convert url_id to int
	number, err := strconv.Atoi(url_id)
	if err != nil {

		c.JSON(400, gin.H{
			"message": "url_id type is not int",
			"data":    nil,
		})

		return
	}

	url = services.GetUrl(number)

	c.JSON(200, gin.H{
		"message": "",
		"data":    url,
	})
}

func UpdateUrl(c *gin.Context) {
	var urlBody structs.UrlBody

	if err := c.ShouldBindJSON(&urlBody); err != nil {
		c.JSON(422, gin.H{"error": err.Error()})
		return
	}

	url_id := c.Param("url_id")
	// var url models.Url

	// convert url_id to int
	number, err := strconv.Atoi(url_id)
	if err != nil {

		c.JSON(400, gin.H{
			"message": "url_id type is not int",
			"data":    nil,
		})

		return
	}

	url, message := services.UpdateUrl(number, urlBody)

	if url != true {
		c.JSON(400, gin.H{
			"message": message,
			"data":    nil,
		})
	}

	c.JSON(200, gin.H{
		"message": message,
		"data":    nil,
	})
}

func DeleteUrl(c *gin.Context) {

	url_id := c.Param("url_id")

	// convert url_id to int
	number, err := strconv.Atoi(url_id)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "url_id type is not int",
			"data":    nil,
		})
		return
	}

	url, message := services.DeleteUrl(number)

	if url != true {
		c.JSON(400, gin.H{
			"message": message,
			"data":    nil,
		})
	}

	c.JSON(200, gin.H{
		"message": message,
		"data":    nil,
	})
}

func CallbackUrl(c *gin.Context) {
	url := c.Param("url")
	fmt.Println(c.Request.Header["Referer"])
	fmt.Println(c.Request.Header["User-Agent"])
	c.JSON(200, gin.H{
		"message": url,
		"data":    []gin.H{},
	})
}
