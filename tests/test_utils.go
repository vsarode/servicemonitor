package tests

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"servicemonitor/models"
)

func getDb() *gorm.DB {
	database, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		log.Println("Failed to connect to database!")
		panic(err)
	}
	return database
}

func setupSchema(ctx models.Context) {
	ctx.Db.AutoMigrate(&models.ServiceInfo{})
	ctx.Db.AutoMigrate(&models.HelthStat{})
}

func insertData(ctx models.Context) {
	serviceInfo := models.ServiceInfo{
		URL: "http://www.youtube.com",
		ServiceName:  "youtube",
		ResponseCode: 200,
	}
	serviceHelth := models.HelthStat{
		ServiceId: 1,
		Status:    true,
		Date:      "sdsadasdas",
		TimeTaken: "2133sec",
	}
	ctx.Db.Create(&serviceInfo)
	ctx.Db.Create(&serviceHelth)
}
