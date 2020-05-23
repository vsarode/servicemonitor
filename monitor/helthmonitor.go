package monitor

import (
	"log"
	"net/http"
	"servicemonitor/models"
	"time"
)

func GetEndPointHealth(endPoint models.ServiceInfo) *models.HelthStat {
	startTime := time.Now()
	resp, err := http.Get(endPoint.URL)
	if err != nil {
		log.Println("Error while getting response for:" + endPoint.ServiceName + " Error:" + err.Error())
		return nil
	}
	status := false
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
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

func GetEndPointsHealth(endPoints []models.ServiceInfo) []*models.HelthStat {
	var results []*models.HelthStat
	for _, e := range endPoints {
		health := GetEndPointHealth(e)
		if health == nil {
			log.Println("Error in getendpointhelth for:", e)
		} else {
			results = append(results, health)
		}
	}
	return results
}

func RecordHealth(ctx models.Context) {
	var endpoints []models.ServiceInfo
	ctx.Db.Find(&endpoints)
	healthRecords := GetEndPointsHealth(endpoints)
	for _, record := range healthRecords {
		ctx.Db.Create(record)
	}
}
