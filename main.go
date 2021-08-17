package main

import (
	"fmt"
	"runtime/debug"
	router "url-shortener/Router"

	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
)

func TryCatch(f func()) func() error {
	return func() (err error) {
		defer func() {
			if panicInfo := recover(); panicInfo != nil {
				err = fmt.Errorf("%v, %s", panicInfo, string(debug.Stack()))
				return
			}
		}()
		f() // calling the decorated function
		return err
	}
}

func main() {

	r := gin.Default()
	//new template engine
	r.HTMLRender = gintemplate.New(gintemplate.TemplateConfig{
		Root:         "views",
		Extension:    ".html",
		Master:       "index",
		Partials:     []string{"partials/ad"},
		DisableCache: true,
	})

	// register home route
	router.RegisterAuthRoute(r)
	router.RegisterUrlsRoute(r)

	r.Run(":8000")
}
