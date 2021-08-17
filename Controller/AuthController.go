package Controller

import (
	services "url-shortener/Services"
	structs "url-shortener/Structs"

	"github.com/gin-gonic/gin"
)

type Formbody struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}

func Login(c *gin.Context) {
	var loginBody structs.LoginBody

	if err := c.ShouldBindJSON(&loginBody); err != nil {
		c.JSON(422, gin.H{"error": err.Error()})
		return
	}

	// c.BindJSON(&loginBody)

	// fmt.Println("user: ", user)
	status, message, response := services.CheckLoginAttempt(loginBody)
	// if no user found
	if status == false {
		c.JSON(401, gin.H{
			"message": message,
			"data": gin.H{
				"name":  response.Name,
				"email": response.Email,
				"token": response.Token,
			},
		})
		return
	}

	c.JSON(200, gin.H{
		"message": message,
		"data": gin.H{
			"name":  response.Name,
			"email": response.Email,
			"token": response.Token,
		},
	})
}

func Register(c *gin.Context) {
	var registerBody structs.RegisterBody

	if err := c.ShouldBindJSON(&registerBody); err != nil {
		c.JSON(422, gin.H{"error": err.Error()})
		return
	}

	// c.BindJSON(&registerBody)

	// var user models.User

	status, message, response := services.StoreUser(registerBody)

	if status == false {

		c.JSON(403, gin.H{
			"message": message,
			"data": gin.H{
				"id":    0,
				"email": "",
				"name":  "",
				"token": "",
			},
		})

		return
	}

	// fmt.Println("user: ", user)
	// fmt.Println("response: ", response)
	// if no user found

	c.JSON(200, gin.H{
		"message": message,
		"data": gin.H{
			"id":    response.ID,
			"email": response.Email,
			"name":  response.Name,
			"token": response.Token,
		},
	})
}
