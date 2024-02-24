package entity

import (
	"caiyu-go/dao"
	"caiyu-go/dao/models"
	"caiyu-go/system_common"
	"time"
)

type DailyAstroEntity struct {
	Astro []struct {
		Date    string `json:"date"`
		Sunrise struct {
			Time string `json:"time"`
		} `json:"sunrise"`
		Sunset struct {
			Time string `json:"time"`
		} `json:"sunset"`
	} `json:"astro"`
}

func (d *DailyAstroEntity) ConvToModel() (result []models.DailyAstro) {

	for i := range d.Astro {
		element := d.Astro[i]
		elementTime, _ := time.Parse(system_common.Layout, element.Date)
		sunrise, _ := time.Parse("15:04", element.Sunrise.Time)
		sunset, _ := time.Parse("15:04", element.Sunset.Time)

		year, month, day := elementTime.Date()
		dateSunrise := time.Date(year, month, day, sunrise.Hour(), sunrise.Minute(), sunrise.Second(), sunrise.Nanosecond(), sunrise.Location())
		dateSunset := time.Date(year, month, day, sunset.Hour(), sunset.Minute(), sunset.Second(), sunset.Nanosecond(), sunset.Location())

		result = append(result, models.DailyAstro{
			Date:        elementTime,
			SunriseTime: &(dateSunrise),
			SunsetTime:  &(dateSunset),
		})

	}

	return result
}

type DailyPrecipitation08H20HEntity struct {
	Precipitation08H20H []struct {
		Date        string  `json:"date"`
		Max         float64 `json:"max"`
		Min         float64 `json:"min"`
		Avg         float64 `json:"avg"`
		Probability float64 `json:"probability"`
	} `json:"precipitation_08h_20h"`
}

func (d *DailyPrecipitation08H20HEntity) ConvToModel() (result []models.DailyPrecipitation08h20h) {
	for i := range d.Precipitation08H20H {
		element := d.Precipitation08H20H[i]

		elementTime, _ := time.Parse(system_common.Layout, element.Date)
		result = append(result, models.DailyPrecipitation08h20h{
			Date:        elementTime,
			Max:         &element.Max,
			Min:         &element.Min,
			Avg:         &element.Avg,
			Probability: &element.Probability,
		})
	}

	return result
}

type DailyPrecipitation20H32HEntity struct {
	Precipitation20H32H []struct {
		Date        string  `json:"date"`
		Max         float64 `json:"max"`
		Min         float64 `json:"min"`
		Avg         float64 `json:"avg"`
		Probability float64 `json:"probability"`
	} `json:"precipitation_20h_32h"`
}

func (d *DailyPrecipitation20H32HEntity) ConvToModel() (result []models.DailyPrecipitation20h32h) {
	for i := range d.Precipitation20H32H {
		element := d.Precipitation20H32H[i]

		elementTime, _ := time.Parse(system_common.Layout, element.Date)
		result = append(result, models.DailyPrecipitation20h32h{
			Date:        elementTime,
			Max:         &element.Max,
			Min:         &element.Min,
			Avg:         &element.Avg,
			Probability: &element.Probability,
		})
	}

	return result
}

type DailyPrecipitationEntity struct {
	Precipitation []struct {
		Date        string  `json:"date"`
		Max         float64 `json:"max"`
		Min         float64 `json:"min"`
		Avg         float64 `json:"avg"`
		Probability float64 `json:"probability"`
	} `json:"precipitation"`
}

func (d *DailyPrecipitationEntity) ConvToModel() (result []models.Precipitation) {
	for i := range d.Precipitation {
		element := d.Precipitation[i]

		elementTime, _ := time.Parse(system_common.Layout, element.Date)
		result = append(result, models.Precipitation{
			UpdateFrequency: daily,
			Date:            elementTime,
			Max:             &element.Max,
			Min:             &element.Min,
			Avg:             &element.Avg,
			Probability:     &element.Probability,
		})
	}

	return result
}

//func (d *DailyPrecipitationEntity) Alert() (alerts []WeatherAlert) {
//	precipitations := d.ConvToModel()
//	NewAlertWrapper(precipitations).Inject(
//		AVal("Probability").BiggerThan(0.6).Then(func(entity interface{}) interface{} {
//			entityValue := entity.(models.Precipitation)
//
//			alerts = append(alerts, WeatherAlert{Description: fmt.Sprintf("在 %s 【%s】有 %d%% 的概率 降水!",
//				entityValue.Date,
//				time.Now().Sub(entityValue.Date).String(),
//				entityValue.Probability)})
//
//			return nil
//		}),
//	)
//
//	return
//}

type DailyTemperatureEntity struct {
	Temperature []struct {
		Date string  `json:"date"`
		Max  float64 `json:"max"`
		Min  float64 `json:"min"`
		Avg  float64 `json:"avg"`
	} `json:"temperature"`
}

func (t DailyTemperatureEntity) ConvToModel() (result []models.Temperature) {
	for i := range t.Temperature {
		element := t.Temperature[i]

		elementTime, _ := time.Parse(system_common.Layout, element.Date)
		result = append(result, models.Temperature{
			UpdateFrequency: daily,
			Date:            elementTime,
			Max:             &element.Max,
			Min:             &element.Min,
			Avg:             &element.Avg,
		})
	}

	return result
}

type DailyTemperature08H20HEntity struct {
	Temperature08H20H []struct {
		Date string  `json:"date"`
		Max  float64 `json:"max"`
		Min  float64 `json:"min"`
		Avg  float64 `json:"avg"`
	} `json:"temperature_08h_20h"`
}

func (t DailyTemperature08H20HEntity) ConvToModel() (result []models.DailyTemperature08H20H) {
	for i := range t.Temperature08H20H {
		element := t.Temperature08H20H[i]

		elementTime, _ := time.Parse(system_common.Layout, element.Date)
		result = append(result, models.DailyTemperature08H20H{
			Date: elementTime,
			Max:  &element.Max,
			Min:  &element.Min,
			Avg:  &element.Avg,
		})
	}

	return result
}

type DailyTemperature20H32HHEntity struct {
	Temperature20H32H []struct {
		Date string  `json:"date"`
		Max  float64 `json:"max"`
		Min  float64 `json:"min"`
		Avg  float64 `json:"avg"`
	} `json:"temperature_20h_32h"`
}

func (t DailyTemperature20H32HHEntity) ConvToModel() (result []models.DailyTemperature20H32H) {
	for i := range t.Temperature20H32H {
		element := t.Temperature20H32H[i]
		elementTime, _ := time.Parse(system_common.Layout, element.Date)
		result = append(result, models.DailyTemperature20H32H{
			Date: elementTime,
			Max:  &element.Max,
			Min:  &element.Min,
			Avg:  &element.Avg,
		})
	}

	return result
}

type DailyWindEntity struct {
	Wind []struct {
		Date string `json:"date"`
		Max  struct {
			Speed     float64 `json:"speed"`
			Direction float64 `json:"direction"`
		} `json:"max"`
		Min struct {
			Speed     float64 `json:"speed"`
			Direction float64 `json:"direction"`
		} `json:"min"`
		Avg struct {
			Speed     float64 `json:"speed"`
			Direction float64 `json:"direction"`
		} `json:"avg"`
	} `json:"wind"`
}

func (w DailyWindEntity) ConvToModel() (result []models.Wind) {
	for i := range w.Wind {
		element := w.Wind[i]
		elementTime, _ := time.Parse(system_common.Layout, element.Date)
		result = append(result, models.Wind{
			UpdateFrequency: daily,
			Date:            elementTime,
			MaxSpeed:        &element.Max.Speed,
			MaxDirection:    &element.Max.Direction,
			MinSpeed:        &element.Min.Speed,
			MinDirection:    &element.Min.Direction,
			AvgSpeed:        &element.Avg.Speed,
			AvgDirection:    &element.Avg.Direction,
		})
	}

	return result
}

type DailyWind08H20HEntity struct {
	Wind08h20h []struct {
		Date string `json:"date"`
		Max  struct {
			Speed     float64 `json:"speed"`
			Direction float64 `json:"direction"`
		} `json:"max"`
		Min struct {
			Speed     float64 `json:"speed"`
			Direction float64 `json:"direction"`
		} `json:"min"`
		Avg struct {
			Speed     float64 `json:"speed"`
			Direction float64 `json:"direction"`
		} `json:"avg"`
	} `json:"wind_08h_20h"`
}

func (w DailyWind08H20HEntity) ConvToModel() (result []models.DailyWind08h20h) {
	for i := range w.Wind08h20h {
		element := w.Wind08h20h[i]
		elementTime, _ := time.Parse(system_common.Layout, element.Date)
		result = append(result, models.DailyWind08h20h{
			Date:         elementTime,
			MaxSpeed:     &element.Max.Speed,
			MaxDirection: &element.Max.Direction,
			MinSpeed:     &element.Min.Speed,
			MinDirection: &element.Min.Direction,
			AvgSpeed:     &element.Avg.Speed,
			AvgDirection: &element.Avg.Direction,
		})
	}

	return result
}

type DailyWind20H32HEntity struct {
	Wind20H32H []struct {
		Date string `json:"date"`
		Max  struct {
			Speed     float64 `json:"speed"`
			Direction float64 `json:"direction"`
		} `json:"max"`
		Min struct {
			Speed     float64 `json:"speed"`
			Direction float64 `json:"direction"`
		} `json:"min"`
		Avg struct {
			Speed     float64 `json:"speed"`
			Direction float64 `json:"direction"`
		} `json:"avg"`
	} `json:"wind_20h_32h"`
}

func (w DailyWind20H32HEntity) ConvToModel() (result []models.DailyWind20h32h) {
	for i := range w.Wind20H32H {
		element := w.Wind20H32H[i]
		elementTime, _ := time.Parse(system_common.Layout, element.Date)
		result = append(result, models.DailyWind20h32h{
			Date:         elementTime,
			MaxSpeed:     &element.Max.Speed,
			MaxDirection: &element.Max.Direction,
			MinSpeed:     &element.Min.Speed,
			MinDirection: &element.Min.Direction,
			AvgSpeed:     &element.Avg.Speed,
			AvgDirection: &element.Avg.Direction,
		})
	}

	return result
}

// DailyHumidityEntity 地表 2 米相对湿度(%)
type DailyHumidityEntity struct {
	Humidity []struct {
		Date string  `json:"date"`
		Max  float64 `json:"max"`
		Min  float64 `json:"min"`
		Avg  float64 `json:"avg"`
	} `json:"humidity"`
}

func (d DailyHumidityEntity) ConvToModel() (result []models.Humidity) {
	for i := range d.Humidity {
		element := d.Humidity[i]
		elementTime, _ := time.Parse(system_common.Layout, element.Date)
		result = append(result, models.Humidity{
			UpdateFrequency: daily,
			Date:            elementTime,
			Max:             &element.Max,
			Min:             &element.Min,
			Avg:             &element.Avg,
		})

	}

	return result
}

type DailyCloudrateEntity struct {
	Cloudrate []struct {
		Date string  `json:"date"`
		Max  float64 `json:"max"`
		Min  float64 `json:"min"`
		Avg  float64 `json:"avg"`
	} `json:"cloudrate"`
}

func (d DailyCloudrateEntity) ConvToModel() (result []models.CloudRate) {
	for i := range d.Cloudrate {
		element := d.Cloudrate[i]
		elementTime, _ := time.Parse(system_common.Layout, element.Date)
		result = append(result, models.CloudRate{
			UpdateFrequency: daily,
			Date:            elementTime,
			Max:             &element.Max,
			Min:             &element.Min,
			Avg:             &element.Avg,
		})

	}

	return result
}

type DailyPressureEntity struct {
	Pressure []struct {
		Date string  `json:"date"`
		Max  float64 `json:"max"`
		Min  float64 `json:"min"`
		Avg  float64 `json:"avg"`
	} `json:"pressure"`
}

func (d DailyPressureEntity) ConvToModel() (result []models.Pressure) {
	for i := range d.Pressure {
		element := d.Pressure[i]
		elementTime, _ := time.Parse(system_common.Layout, element.Date)
		result = append(result, models.Pressure{
			UpdateFrequency: daily,
			Date:            elementTime,
			Max:             &element.Max,
			Min:             &element.Min,
			Avg:             &element.Avg,
		})

	}

	return result
}

type DailyVisibilityEntity struct {
	Visibility []struct {
		Date string  `json:"date"`
		Max  float64 `json:"max"`
		Min  float64 `json:"min"`
		Avg  float64 `json:"avg"`
	} `json:"visibility"`
}

func (d DailyVisibilityEntity) ConvToModel() (result []models.Visibility) {
	for i := range d.Visibility {
		element := d.Visibility[i]
		elementTime, _ := time.Parse(system_common.Layout, element.Date)
		result = append(result, models.Visibility{
			UpdateFrequency: daily,
			Date:            elementTime,
			Max:             &element.Max,
			Min:             &element.Min,
			Avg:             &element.Avg,
		})

	}

	return result
}

type DailyDswfEntity struct {
	Dswrf []struct {
		Date string  `json:"date"`
		Max  float64 `json:"max"`
		Min  float64 `json:"min"`
		Avg  float64 `json:"avg"`
	} `json:"dswrf"`
}

func (d DailyDswfEntity) ConvToModel() (result []models.DSWRF) {
	for i := range d.Dswrf {
		element := d.Dswrf[i]
		elementTime, _ := time.Parse(system_common.Layout, element.Date)
		result = append(result, models.DSWRF{
			UpdateFrequency: daily,
			Date:            elementTime,
			Max:             &element.Max,
			Min:             &element.Min,
			Avg:             &element.Avg,
		})

	}

	return result
}

type DailyAqiEntity struct {
	Aqi []struct {
		Date string `json:"date"`
		Max  struct {
			Chn float64 `json:"chn"`
			Usa float64 `json:"usa"`
		} `json:"max"`
		Avg struct {
			Chn float64 `json:"chn"`
			Usa float64 `json:"usa"`
		} `json:"avg"`
		Min struct {
			Chn float64 `json:"chn"`
			Usa float64 `json:"usa"`
		} `json:"min"`
	} `json:"aqi"`
}

func (d DailyAqiEntity) ConvToModel() (result []models.AQI) {
	for i := range d.Aqi {
		element := d.Aqi[i]
		elementTime, _ := time.Parse(system_common.Layout, element.Date)
		result = append(result, models.AQI{
			UpdateFrequency: daily,
			Date:            elementTime,
			MaxCHN:          &element.Max.Chn,
			MaxUSA:          &element.Max.Usa,
			AvgCHN:          &element.Avg.Chn,
			AvgUSA:          &element.Avg.Usa,
			MinCHN:          &element.Min.Chn,
			MinUSA:          &element.Min.Usa,
		})

	}

	return result
}

type DailyPm25Entity struct {
	Pm25 []struct {
		Date string  `json:"date"`
		Max  float64 `json:"max"`
		Avg  float64 `json:"avg"`
		Min  float64 `json:"min"`
	} `json:"pm25"`
}

func (d DailyPm25Entity) ConvToModel() (result []models.PM25) {
	for i := range d.Pm25 {
		element := d.Pm25[i]
		elementTime, _ := time.Parse(system_common.Layout, element.Date)
		result = append(result, models.PM25{
			UpdateFrequency: daily,
			Date:            elementTime,
			Max:             &element.Max,
			Min:             &element.Min,
			Avg:             &element.Avg,
		})

	}

	return result
}

type DailySkyConEntity struct {
	Skycon []struct {
		Date  string `json:"date"`
		Value string `json:"value"`
	} `json:"skycon"`
}

func (d DailySkyConEntity) ConvToModel() (result []models.Skycon) {
	for i := range d.Skycon {
		element := d.Skycon[i]
		elementTime, _ := time.Parse(system_common.Layout, element.Date)
		result = append(result, models.Skycon{
			UpdateFrequency: daily,
			Date:            elementTime,
			Value:           &element.Value,
		})

	}

	return result
}

type DailySkyCon08H20HEntity struct {
	Skycon08H20H []struct {
		Date  string `json:"date"`
		Value string `json:"value"`
	} `json:"skycon_08h_20h"`
}

func (d DailySkyCon08H20HEntity) ConvToModel() (result []models.DailySkycon08H20H) {
	for i := range d.Skycon08H20H {
		element := d.Skycon08H20H[i]
		elementTime, _ := time.Parse(system_common.Layout, element.Date)
		result = append(result, models.DailySkycon08H20H{
			Date:  elementTime,
			Value: &element.Value,
		})

	}

	return result
}

type DailySkyCon20H32HEntity struct {
	Skycon20H32H []struct {
		Date  string `json:"date"`
		Value string `json:"value"`
	} `json:"skycon_20h_32h"`
}

func (d DailySkyCon20H32HEntity) ConvToModel() (result []models.DailySkycon20H32H) {
	for i := range d.Skycon20H32H {
		element := d.Skycon20H32H[i]
		elementTime, _ := time.Parse(system_common.Layout, element.Date)
		result = append(result, models.DailySkycon20H32H{
			Date:  elementTime,
			Value: &element.Value,
		})

	}

	return result
}

type DailyUltravioletEntity struct {
	Ultraviolet []struct {
		Date  string `json:"date"`
		Index string `json:"index"`
		Desc  string `json:"desc"`
	} `json:"ultraviolet"`
}

func (d DailyUltravioletEntity) ConvToModel() (result []models.Ultraviolet) {
	for i := range d.Ultraviolet {
		element := d.Ultraviolet[i]
		elementTime, _ := time.Parse(system_common.Layout, element.Date)
		result = append(result, models.Ultraviolet{
			UpdateFrequency: daily,
			Date:            elementTime,
			Index:           &element.Index,
			Description:     &element.Desc,
		})

	}

	return result
}

type DailyCarWashingEntity struct {
	CarWashing []struct {
		Date  string `json:"date"`
		Index string `json:"index"`
		Desc  string `json:"desc"`
	} `json:"carWashing"`
}

func (d DailyCarWashingEntity) ConvToModel() (result []models.DailyCarWashing) {
	for i := range d.CarWashing {
		element := d.CarWashing[i]
		elementTime, _ := time.Parse(system_common.Layout, element.Date)
		result = append(result, models.DailyCarWashing{
			Date:        elementTime,
			Index:       &element.Index,
			Description: &element.Desc,
		})

	}

	return result
}

type DailyDressingEntity struct {
	Dressing []struct {
		Date  string `json:"date"`
		Index string `json:"index"`
		Desc  string `json:"desc"`
	} `json:"dressing"`
}

func (d DailyDressingEntity) ConvToModel() (result []models.DailyDressing) {
	for i := range d.Dressing {
		element := d.Dressing[i]
		elementTime, _ := time.Parse(system_common.Layout, element.Date)
		result = append(result, models.DailyDressing{
			Date:        elementTime,
			Index:       &element.Index,
			Description: &element.Desc,
		})

	}

	return result
}

type DailyComfortEntity struct {
	Comfort []struct {
		Date  string `json:"date"`
		Index string `json:"index"`
		Desc  string `json:"desc"`
	} `json:"comfort"`
}

func (d DailyComfortEntity) ConvToModel() (result []models.Comfort) {
	for i := range d.Comfort {
		element := d.Comfort[i]
		elementTime, _ := time.Parse(system_common.Layout, element.Date)
		result = append(result, models.Comfort{
			UpdateFrequency: daily,
			Date:            elementTime,
			Index:           &element.Index,
			Description:     &element.Desc,
		})

	}

	return result
}

type DailyColdRiskEntity struct {
	ColdRisk []struct {
		Date  string `json:"date"`
		Index string `json:"index"`
		Desc  string `json:"desc"`
	} `json:"coldRisk"`
}

func (d DailyColdRiskEntity) ConvToModel() (result []models.DailyColdRisk) {
	for i := range d.ColdRisk {
		element := d.ColdRisk[i]
		elementTime, _ := time.Parse(system_common.Layout, element.Date)
		result = append(result, models.DailyColdRisk{
			Date:        elementTime,
			Index:       &element.Index,
			Description: &element.Desc,
		})

	}

	return result
}

type DailyEntity struct {
	Daily struct {
		Status string `json:"status"`
		DailyAstroEntity
		DailyPrecipitation08H20HEntity
		DailyPrecipitation20H32HEntity
		DailyPrecipitationEntity
		DailyTemperatureEntity
		DailyTemperature08H20HEntity
		DailyTemperature20H32HHEntity
		DailyWindEntity
		DailyWind08H20HEntity
		DailyWind20H32HEntity
		DailyHumidityEntity
		DailyCloudrateEntity
		DailyPressureEntity
		DailyVisibilityEntity
		DailyDswfEntity
		AirQuality struct {
			DailyAqiEntity
			DailyPm25Entity
		} `json:"air_quality"`
		DailySkyConEntity
		DailySkyCon08H20HEntity
		DailySkyCon20H32HEntity
		LifeIndex struct {
			DailyUltravioletEntity
			DailyCarWashingEntity
			DailyDressingEntity
			DailyComfortEntity
			DailyColdRiskEntity
		} `json:"life_index"`
	} `json:"daily"`
}

func (d *DailyEntity) Save() {
	// 尝试保存到数据库
	astros := d.Daily.DailyAstroEntity.ConvToModel()
	precipitations := d.Daily.DailyPrecipitationEntity.ConvToModel()
	precipitation08h20hs := d.Daily.DailyPrecipitation08H20HEntity.ConvToModel()
	precipitation20h32hs := d.Daily.DailyPrecipitation20H32HEntity.ConvToModel()
	temperatures := d.Daily.DailyTemperatureEntity.ConvToModel()
	temperature08H20HS := d.Daily.DailyTemperature08H20HEntity.ConvToModel()
	temperature20H32HS := d.Daily.DailyTemperature20H32HHEntity.ConvToModel()
	winds := d.Daily.DailyWindEntity.ConvToModel()
	wind08h20hs := d.Daily.DailyWind08H20HEntity.ConvToModel()
	wind20h32hs := d.Daily.DailyWind20H32HEntity.ConvToModel()
	humidities := d.Daily.DailyHumidityEntity.ConvToModel()
	cloudRates := d.Daily.DailyCloudrateEntity.ConvToModel()
	pressures := d.Daily.DailyPressureEntity.ConvToModel()
	visibilities := d.Daily.DailyVisibilityEntity.ConvToModel()
	dswrfs := d.Daily.DailyDswfEntity.ConvToModel()
	aqis := d.Daily.AirQuality.DailyAqiEntity.ConvToModel()
	pm25s := d.Daily.AirQuality.DailyPm25Entity.ConvToModel()
	skycons := d.Daily.DailySkyConEntity.ConvToModel()
	skycon08H20HS := d.Daily.DailySkyCon08H20HEntity.ConvToModel()
	skycon20H32HS := d.Daily.DailySkyCon20H32HEntity.ConvToModel()
	ultraviolets := d.Daily.LifeIndex.DailyUltravioletEntity.ConvToModel()
	carWashings := d.Daily.LifeIndex.DailyCarWashingEntity.ConvToModel()
	dressings := d.Daily.LifeIndex.DailyDressingEntity.ConvToModel()
	comforts := d.Daily.LifeIndex.DailyComfortEntity.ConvToModel()
	coldRisks := d.Daily.LifeIndex.DailyColdRiskEntity.ConvToModel()

	dao.DB.Create(&astros)
	dao.DB.Create(&precipitations)
	dao.DB.Create(&precipitation08h20hs)
	dao.DB.Create(&precipitation20h32hs)
	dao.DB.Create(&temperatures)
	dao.DB.Create(&temperature08H20HS)
	dao.DB.Create(&temperature20H32HS)
	dao.DB.Create(&winds)
	dao.DB.Create(&wind08h20hs)
	dao.DB.Create(&wind20h32hs)
	dao.DB.Create(&humidities)
	dao.DB.Create(&cloudRates)
	dao.DB.Create(&pressures)
	dao.DB.Create(&visibilities)
	dao.DB.Create(&dswrfs)
	dao.DB.Create(&aqis)
	dao.DB.Create(&pm25s)
	dao.DB.Create(&skycons)
	dao.DB.Create(&skycon08H20HS)
	dao.DB.Create(&skycon20H32HS)
	dao.DB.Create(&ultraviolets)
	dao.DB.Create(&carWashings)
	dao.DB.Create(&dressings)
	dao.DB.Create(&comforts)
	dao.DB.Create(&coldRisks)
}
