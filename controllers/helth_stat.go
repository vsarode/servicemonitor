package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"servicemonitor/models"
)

func GetHelthStats(ctx models.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		var serviceInfos []models.ServiceInfo
		ctx.Db.Find(&serviceInfos)

		var helthStats []models.ServiceHelthResponse
		for _, serviceInfo := range serviceInfos {
			var stat models.HelthStat
			ctx.Db.Where("service_id = ?", serviceInfo.ID).Last(&stat)
			helthStats = append(helthStats, models.ServiceHelthResponse{
				ServiceInfo: serviceInfo,
				HelthStat:   &stat,
			})
		}
		c.JSON(http.StatusOK, gin.H{"data": helthStats})
	}
}
