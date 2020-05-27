package main

import (
	"github.com/jasonlvhit/gocron"
	"log"
	"os"
	"servicemonitor/db"
	"servicemonitor/models"
	"servicemonitor/monitor"
	router2 "servicemonitor/router"
	"strconv"
)

func main() {
	database := db.GetDb()
	// Connect to database
	context := models.Context{
		Db: database,
	}
	defer database.Close()
	cronRerunTimeInSeconds, err := strconv.ParseInt(os.Getenv("cron"), 10, 64)
	if err != nil {
		log.Println("Error while geting value from env! Error:", err.Error())
		panic(err)
	}
	db.SetupSchema(context)
	db.InsertDummyEndPoints(context)
	monitor.RecordHealth(context)
	gocron.Every(uint64(cronRerunTimeInSeconds)).Second().Do(monitor.RecordHealth, context)
	gocron.Start()
	router := router2.SetupRouter(context)
	// Run the server
	router.Run()
}
