package main

import (
	"encoding/json"
	"github.com/jasonlvhit/gocron"
	"log"
	"os"
	"servicemonitor/db"
	"servicemonitor/models"
	"servicemonitor/monitor"
	router2 "servicemonitor/router"
)

func main() {
	database := db.GetDb()
	// Connect to database
	context := models.Context{
		Db: database,
	}
	defer database.Close()
	file, err := os.Open("config/local.json")
	if err != nil {
		log.Println("Error while open file! Error:", err.Error())
		panic(err)
	}
	decoder := json.NewDecoder(file)
	config := models.Config{}
	err = decoder.Decode(&config)
	if err != nil {
		log.Println("Error while reading config! Error:", err.Error())
		panic(err)
	}
	db.SetupSchema(context)
	db.InsertDummyEndPoints(context)
	monitor.RecordHealth(context)
	gocron.Every(config.CronRerunTimeInSecond).Second().Do(monitor.RecordHealth, context)
	gocron.Start()
	router := router2.SetupRouter(context)
	// Run the server
	router.Run()
}
