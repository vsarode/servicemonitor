package router


import (
	"github.com/gin-gonic/gin"
	"servicemonitor/controllers"
	"servicemonitor/models"
)

func SetupRouter(ctx models.Context) *gin.Engine {
	router := gin.Default()
	router.Use(controllers.GetAssets())
	apiRoutes := router.Group("/monitor")
	// Routes
	apiRoutes.GET("/service", controllers.GetServiceInfos(ctx))
	apiRoutes.POST("/service", controllers.CreateServiceInfo(ctx))
	apiRoutes.GET("/servicestat", controllers.GetHelthStats(ctx))
	return router
}