package entity

import (
	"caiyu-go/dao"
	"caiyu-go/dao/models"
	"time"
)

type AlertEntity struct {
	Alert struct {
		Status  string `json:"status"`
		Content []struct {
			Province      string    `json:"province"`
			Status        string    `json:"status"`
			Code          string    `json:"code"`
			Description   string    `json:"description"`
			RegionId      string    `json:"regionId"`
			County        string    `json:"county"`
			Pubtimestamp  int       `json:"pubtimestamp"`
			Latlon        []float64 `json:"latlon"`
			City          string    `json:"city"`
			AlertId       string    `json:"alertId"`
			Title         string    `json:"title"`
			Adcode        string    `json:"adcode"`
			Source        string    `json:"source"`
			Location      string    `json:"location"`
			RequestStatus string    `json:"request_status"`
		} `json:"content"`
		Adcodes []struct {
			Adcode int    `json:"adcode"`
			Name   string `json:"name"`
		} `json:"adcodes"`
	} `json:"alert_util"`
}

func (a *AlertEntity) Save() {
	weatherAlerts := a.ConvToModel()
	dao.DB.Create(&weatherAlerts)
}

// ConvToModel 将AlertEntity转换为WeatherAlert模型的切片
func (a *AlertEntity) ConvToModel() []models.WeatherAlert {
	var weatherAlerts []models.WeatherAlert

	for _, content := range a.Alert.Content {
		weatherAlert := models.WeatherAlert{
			Province:      content.Province,
			City:          content.City,
			County:        content.County,
			Status:        content.Status,
			Code:          content.Code,
			Description:   content.Description,
			PubTimestamp:  time.Unix(int64(content.Pubtimestamp), 0), // 将int转换为time.Time
			AlertID:       content.AlertId,
			Title:         content.Title,
			Adcode:        content.Adcode,
			Source:        content.Source,
			Location:      content.Location,
			RequestStatus: content.RequestStatus,
		}

		// 确保Latlon字段有两个元素：纬度和经度
		if len(content.Latlon) == 2 {
			weatherAlert.Latitude = content.Latlon[0]
			weatherAlert.Longitude = content.Latlon[1]
		}

		weatherAlerts = append(weatherAlerts, weatherAlert)
	}

	return weatherAlerts
}
