package monitor

import (
	"log"
	"net/http"
	"servicemonitor/models"
	"time"
)

func getServiceHealth(endPoint models.ServiceInfo) *models.HelthStat {
	startTime := time.Now()
	resp, err := http.Get(endPoint.URL)
	if err != nil {
		log.Println("Error while getting response for:" + endPoint.ServiceName + " Error:" + err.Error())
		return nil
	}
	status := false
	if resp.StatusCode == endPoint.ResponseCode {
		status = true
	}
	health := &models.HelthStat{
		ServiceId: endPoint.ID,
		Status:    status,
		TimeTaken: time.Since(startTime).String(),
		Date:      startTime.Format(time.RFC850),
	}
	return health
}

func getAllServicesHealth(endPoints []models.ServiceInfo) []*models.HelthStat {
	var serviceHelths []*models.HelthStat
	for _, e := range endPoints {
		health := getServiceHealth(e)
		if health == nil {
			log.Println("Error in getServiceHealth for:", e)
		} else {
			serviceHelths = append(serviceHelths, health)
		}
	}
	return serviceHelths
}

func RecordHealth(ctx models.Context) {
	var endpoints []models.ServiceInfo
	ctx.Db.Find(&endpoints)
	healthRecords := getAllServicesHealth(endpoints)
	for _, record := range healthRecords {
		ctx.Db.Create(record)
	}
}
