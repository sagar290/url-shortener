package Model

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var (
	Db         *gorm.DB
	Connection string
	Host       string
	Port       string
	Database   string
	Username   string
	Password   string
	Result     string
)

func init() {
	viper.SetEnvPrefix("url-shortener")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/url-shortener/")
	viper.AddConfigPath("$HOME/.url-shortener")
	viper.AddConfigPath("../")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	Connection := viper.GetString("DB_CONNECTION")
	Host := viper.GetString("DB_HOST")
	Port := viper.GetString("DB_PORT")
	Database := viper.GetString("DB_DATABASE")
	Username := viper.GetString("DB_USERNAME")
	Password := viper.GetString("DB_PASSWORD")

	if Connection == "mysql" {
		dsn := Username + ":" + Password + "@tcp(" + Host + ":" + Port + ")/" + Database + "?charset=utf8mb4&parseTime=True&loc=Local"
		Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		_ = Db
	}

	if Connection == "postgres" {
		dsn := "host=" + Host + " user=" + Username + " password=" + Password + " dbname=" + Database + " port=" + Port + " sslmode=disable TimeZone=Asia/Dhaka"
		Db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		_ = Db
	}

	if Connection == "sqlite" {
		dsn := "sqlserver://" + Username + ":" + Password + "@" + Host + ":" + Port + "?database=" + Database + ""
		Db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		_ = Db
	}

	autoMigrate()
}

// auto migrate table
func autoMigrate() {
	Db.AutoMigrate(Click{})
	Db.AutoMigrate(Url{})
	Db.AutoMigrate(User{})
}
