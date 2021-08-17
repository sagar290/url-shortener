package Middleware

import (
	"fmt"
	"strings"
	services "url-shortener/Services"

	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {

	return func(c *gin.Context) {
		var header []string

		// header = strings.Split(c.Request.Header["Authorization"][0], " ")
		header = c.Request.Header["Authorization"]

		// fmt.Println(header)
		// fmt.Println(len(header))

		// if len(header) == 1 || len(header) == 0 {
		if len(header) == 0 {
			fmt.Println(header)
			c.AbortWithStatusJSON(401, gin.H{
				"message": "unauthenticated",
				"data":    []gin.H{},
			})

			return
		}

		token := strings.Split(header[0], " ")

		if len(token) == 1 || len(token) == 0 {
			c.AbortWithStatusJSON(401, gin.H{
				"message": "unauthenticated",
				"data":    []gin.H{},
			})

			return
		}

		if token[1] == "" {
			c.AbortWithStatusJSON(401, gin.H{
				"message": "unauthenticated",
				"data":    []gin.H{},
			})

			return
		}

		verify := services.VerifyToken(token[1])

		if verify == false {
			c.AbortWithStatusJSON(401, gin.H{
				"message": "unauthenticated",
				"data":    []gin.H{},
			})

			return
		}

		c.Next()

	}
}
