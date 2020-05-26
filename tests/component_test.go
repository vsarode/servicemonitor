package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"servicemonitor/models"
	router2 "servicemonitor/router"
	"testing"
)

type TestConfig struct {
	Context models.Context
	router  *gin.Engine
}

var GlobalTestConfig TestConfig

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	database := getDb()
	GlobalTestConfig.Context.Db = database
	setupSchema(GlobalTestConfig.Context)
	insertData(GlobalTestConfig.Context)
	//monitor.RecordHealth(GlobalTestConfig.Context)
	GlobalTestConfig.router = router2.SetupRouter(GlobalTestConfig.Context)
	fmt.Printf("\033[1;36m%s\033[0m", "> Setup completed\n")
}

func teardown() {
	GlobalTestConfig.Context.Db.Close()
	err := os.Remove("test.db")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("\033[1;36m%s\033[0m", "> Teardown completed")
	fmt.Printf("\n")
}

func TestServiceStatAPIGETSuccess(t *testing.T) {
	expectedResponse := models.ServiceHelthResponse{
		ServiceInfo: models.ServiceInfo{
			ID:           0,
			URL:          "http://www.youtube.com",
			ServiceName:  "youtube",
			ResponseCode: 200,
		},
		HelthStat: &models.HelthStat{
			ID:        0,
			ServiceId: 0,
			Status:    true,
		},
	}
	type Res struct {
		Data []models.ServiceHelthResponse `json:"data"`
	}
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/monitor/servicestat", nil)
	GlobalTestConfig.router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	response := Res{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Error(err.Error())
	}
	assert.Equal(t, expectedResponse.HelthStat.Status, response.Data[0].HelthStat.Status)
	assert.Equal(t, expectedResponse.ServiceInfo.ServiceName, response.Data[0].ServiceInfo.ServiceName)
}


func TestServiceAPIPostSuccess(t *testing.T) {
	w := httptest.NewRecorder()
	body := models.ServiceInfo{
		ID:           0,
		URL:          "http://www.google.com",
		ServiceName:  "Google",
		ResponseCode: 200,
	}
	requestBody, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/monitor/service", bytes.NewReader(requestBody))
	GlobalTestConfig.router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestServiceAPIGETSuccess(t *testing.T) {
	expectedResponse := models.ServiceInfo{
		ID:           1,
		URL:          "http://www.youtube.com",
		ServiceName:  "youtube",
		ResponseCode: 200,
	}
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/monitor/service", nil)
	GlobalTestConfig.router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	type Res struct {
		Data []models.ServiceInfo `json:"data"`
	}
	response := Res{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, expectedResponse, response.Data[0])
}
