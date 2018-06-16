package main

import (
	"fmt"
	"github.com/dangkaka/leaderboard/scores"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
	"github.com/zalora/tempo/helper"
)

func main() {
	viper.SetConfigType("json")
	viper.SetConfigFile("config/config.json")
	err := viper.ReadInConfig()
	if err != nil {
		helper.ReportFatal("Failed to read config: %v", err)
		return
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
		helper.ReportFatal("Something went horribly wrong: %v", err)
		return
	}
	defer db.Close()
	db.LogMode(true)
	//schema migrations
	db.AutoMigrate(&scores.Score{})

	r := NewRouter(db)
	r.Run(viper.GetString("server.address"))
}
