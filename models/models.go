package models

import (
	"github.com/jinzhu/gorm"
)

//ServiceInfo ...
type ServiceInfo struct {
	ID           uint   `json:"id" gorm:"primary_key"`
	URL          string `json:"url" gorm:"unique;not null"`
	ServiceName  string `json: "serviceName"`
	ResponseCode int    `json: "responseCode"`
}

//HelthStat ...
type HelthStat struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	ServiceId uint   `json:"serviceId"`
	Status    bool   `json:"status"`
	Date      string `json:"date"`
	TimeTaken string `json:"timeTaken"`
}

//ServiceHelthResponse ...
type ServiceHelthResponse struct {
	ServiceInfo ServiceInfo `json: "service"`
	HelthStat   *HelthStat  `json: "helthStatus"`
}

type Context struct {
	Db *gorm.DB
}

type Config struct {
	CronRerunTimeInSecond uint64
}
