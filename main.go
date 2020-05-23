package main

import (
	"servicemonitor/controllers"
	"servicemonitor/db"
	"servicemonitor/models"

	"github.com/gin-gonic/gin"
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
	apiRoutes := r.Group("/monitor")
	// Routes
	apiRoutes.GET("/service", controllers.GetServiceInfos(context))
	apiRoutes.GET("/service/:id", controllers.GetServiceInfo(context))
	apiRoutes.POST("/service", controllers.CreateServiceInfo(context))
	apiRoutes.GET("/servicestat", controllers.GetHelthStats(context))

	// Run the server
	r.Run()
}
