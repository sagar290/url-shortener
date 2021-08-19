package Controller

import (
	"net/http"
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
	slug := c.Param("slug")
	// query := c.Request.URL.Query()
	var referer []string
	var refer_type string
	referer = c.Request.Header["Referer"]
	// get referer
	if len(referer) == 0 {
		referer = c.Request.Header["Referrer"]
		if len(referer) == 0 {
			refer_type = "unknown"
		}
	}

	if refer_type == "" {
		refer_type = referer[0]
	}

	url := services.GetUrlBySlug(slug)
	// fmt.Println(slug)
	if url.Url_id == 0 {
		c.JSON(404, gin.H{
			"message": "no url found",
			"data":    []gin.H{},
		})

		return
	}

	services.AddClick(int(url.Url_id), refer_type)

	c.Redirect(http.StatusMovedPermanently, url.Redirect_url)
}
