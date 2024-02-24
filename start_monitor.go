package main

import (
	"caiyu-go/api"
	"caiyu-go/dao"
	"caiyu-go/dao/models"
	"time"
)

//func MonitorDaily

func main() {
	// 每分钟读取监控位置，对得到的位置列表进行监控
	ticker := time.NewTicker(time.Second)

	locations := make([]models.MonitoringLocation, 0)
	for range ticker.C {
		dao.DB.Find(&locations)
		for _, location := range locations {
			entity := api.GetWeather([]float64{location.Longitude, location.Latitude}, 15, 360)
			entity.Save()

			// 预警系统分两部分，已发生事件解读/目前正发生事件解读/未发生事件预测

			//entity.DailyEntity.Daily.DailyPressureEntity
			// 监控预警信息
			//DailyAstroEntity # 不进行监控
			//DailyPrecipitation08H20HEntity
			//DailyPrecipitation20H32HEntity
			//DailyPrecipitationEntity # 监控降水量
			//DailyTemperatureEntity
			//DailyTemperature08H20HEntity
			//DailyTemperature20H32HHEntity
			//DailyWindEntity
			//DailyWind08H20HEntity
			//DailyWind20H32HEntity
			//DailyHumidityEntity
			//DailyCloudrateEntity
			//DailyPressureEntity
			//DailyVisibilityEntity
			//DailyDswfEntity
			//AirQuality struct {
			//	DailyAqiEntity
			//	DailyPm25Entity
			//} `json:"air_quality"`
			//DailySkyConEntity
			//DailySkyCon08H20HEntity
			//DailySkyCon20H32HEntity
			//LifeIndex struct {
			//	DailyUltravioletEntity
			//	DailyCarWashingEntity
			//	DailyDressingEntity
			//	DailyComfortEntity
			//	DailyColdRiskEntity

		}

		ticker.Reset(time.Hour * 2)
	}
}
