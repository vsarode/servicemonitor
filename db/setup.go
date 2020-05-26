package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"servicemonitor/models"
)

func GetDb() *gorm.DB {
	database, err := gorm.Open("sqlite3", "monitor.db")
	if err != nil {
		println("Failed to connect to database!")
		panic(err)
	}
	return database
}

func SetupSchema(ctx models.Context) {
	ctx.Db.AutoMigrate(&models.ServiceInfo{})
	ctx.Db.AutoMigrate(&models.HelthStat{})
}

func InsertDummyEndPoints(context models.Context) {
	endPoints := []models.ServiceInfo{
		{
			URL: "https://golang.org/pkg/time/#example_Time_Format",
			ServiceName: "Golang",
			ResponseCode: 200,
		},
		{
			URL: "https://www.youtube.com/",
			ServiceName: "Youtube",
			ResponseCode: 200,
		},
	}
	for _, endPoint := range endPoints {
		context.Db.Create(&endPoint)
	}
}