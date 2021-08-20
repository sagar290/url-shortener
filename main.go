package main

import (
	"fmt"
	"runtime/debug"
	"strings"
	models "url-shortener/Model"
	router "url-shortener/Router"

	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
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

func init() {
	viper.SetEnvPrefix("url-shortener")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/url-shortener/")
	viper.AddConfigPath("$HOME/.url-shortener")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	// initialize database
	models.Init()
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
	router.RegisterFrontendRoute(r)

	r.Run(":8000")
}
