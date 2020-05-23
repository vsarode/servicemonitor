package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jasonlvhit/gocron"
	"servicemonitor/controllers"
	"servicemonitor/db"
	"servicemonitor/models"
	"servicemonitor/monitor"
)

func main() {
	database := db.GetDb()
	// Connect to database
	context := models.Context{
		Db: database,
	}
	db.SetupSchema(context)
	db.InsertDummyEndPoints(context)

	r := gin.Default()
	monitor.RecordHealth(context)
	gocron.Every(5).Second().Do(monitor.RecordHealth, context)
	gocron.Start()
	apiRoutes := r.Group("/monitor")
	// Routes
	apiRoutes.GET("/service", controllers.GetServiceInfos(context))
	apiRoutes.GET("/service/:id", controllers.GetServiceInfo(context))
	apiRoutes.POST("/service", controllers.CreateServiceInfo(context))
	apiRoutes.GET("/servicestat", controllers.GetHelthStats(context))

	// Run the server
	r.Run()
}
