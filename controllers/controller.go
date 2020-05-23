package controllers

import (
	"net/http"
	"servicemonitor/models"
	"servicemonitor/monitor"

	"github.com/gin-gonic/gin"
)

//HelthStat ..
type HelthStat struct {
	ServiceId uint   `json:"serviceId"`
	Status    bool   `json:"status" binding:"required"`
	Date      string `json:"requestedTime" binding:"required"`
	TimeTaken string `json:"timeTaken" binding:"required"`
}

//ServiceInfo ...
type ServiceInfo struct {
	URL         string `json:"url" binding:"required"`
	ServiceName string `json:"serviceName" binding:"required"`
}

//GetServiceInfos ...
func GetServiceInfos(ctx models.Context) func(c *gin.Context) {
	return func(c *gin.Context) {

		var urls []models.ServiceInfo
		ctx.Db.Find(&urls)

		c.JSON(http.StatusOK, gin.H{"data": urls})
	}
}

// GET /stats
// Find all stats
func GetHelthStats(ctx models.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		var serviceInfos []models.ServiceInfo
		ctx.Db.Find(&serviceInfos)

		var stats []models.ServiceHelthResponse
		for _, serviceInfo := range serviceInfos {
			var stat models.HelthStat
			ctx.Db.Where("service_id = ?", serviceInfo.ID).Last(&stat)
			stats = append(stats, models.ServiceHelthResponse{
				serviceInfo,
				&stat,
			})
		}

		c.JSON(http.StatusOK, gin.H{"data": stats})
	}
}

// POST /stat
// Create new stat
func CreateHelthStat(ctx models.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		// Validate input
		var input HelthStat
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// Create stat
		stat := models.HelthStat{ServiceId: input.ServiceId, Status: input.Status, Date: input.Date, TimeTaken: input.TimeTaken}
		ctx.Db.Create(&stat)

		c.JSON(http.StatusOK, gin.H{"data": stat})
	}
}

// Create new URL
func CreateServiceInfo(ctx models.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		// Validate input
		var input ServiceInfo
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// Create stat
		serviceInfo := models.ServiceInfo{URL: input.URL, ServiceName: input.ServiceName}
		ctx.Db.Create(&serviceInfo)
		c.JSON(http.StatusOK, gin.H{"data": serviceInfo})
		monitor.RecordHealth(ctx)
	}
}
