package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"net/http"
	"github.com/dangkaka/leaderboard/scores"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigType("json")
	viper.SetConfigFile("config/config.json")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	dsn := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		dbName)
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	db.AutoMigrate(&scores.Scores{})
	r := gin.Default()
	r.GET("/scores", func(c *gin.Context) {
		c.String(http.StatusOK, "Get scores")
	})
	r.POST("/scores", func(c *gin.Context) {
		c.String(http.StatusOK, "Post scores")
	})
	r.DELETE("/scores/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.String(http.StatusOK, "Delete %s", id)
	})
	r.Run(viper.GetString("server.address"))
}
