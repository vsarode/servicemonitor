package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"servicemonitor/models"
	"servicemonitor/monitor"
)

func GetServiceInfos(ctx models.Context) func(c *gin.Context) {
	return func(c *gin.Context) {

		var serviceInfos []models.ServiceInfo
		ctx.Db.Find(&serviceInfos)

		c.JSON(http.StatusOK, gin.H{"data": serviceInfos})
	}
}

func CreateServiceInfo(ctx models.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		var serviceInfo models.ServiceInfo
		if err := c.ShouldBindJSON(&serviceInfo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.Db.Create(&serviceInfo)
		c.JSON(http.StatusOK, gin.H{"data": serviceInfo})
		monitor.RecordHealth(ctx)
	}
}
