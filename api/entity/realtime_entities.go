package entity

import (
	"caiyu-go/dao"
	"caiyu-go/dao/models"
	"fmt"
	"time"
)

type RealtimePrecipitationEntity struct {
	Precipitation struct {
		Local struct {
			Status     string  `json:"status"`
			Datasource string  `json:"datasource"`
			Intensity  float64 `json:"intensity"` // 降水强度
		} `json:"local"`
		Nearest struct {
			Status    string  `json:"status"`
			Distance  float64 `json:"distance"`
			Intensity float64 `json:"intensity"`
		} `json:"nearest"`
	} `json:"precipitation"`
}

func (d *RealtimePrecipitationEntity) ConvToModel() (result models.Precipitation, localResult models.RealTimePrecipitation) {
	element := d.Precipitation

	elementTime := time.Now()

	localResult = models.RealTimePrecipitation{
		Date:             elementTime,
		LocalIntensity:   &element.Local.Intensity,
		LocalDataSource:  &element.Local.Datasource,
		NearestDistance:  &element.Nearest.Distance,
		NearestIntensity: &element.Nearest.Intensity,
	}
	result = models.Precipitation{
		UpdateFrequency: realtime,
		Date:            elementTime,
		Avg:             &element.Local.Intensity,
	}
	return result, localResult
}

type RealtimeTemperatureEntity struct {
	Temperature float64 `json:"temperature"`
}

func (t RealtimeTemperatureEntity) ConvToModel() (result models.Temperature) {

	result = models.Temperature{
		UpdateFrequency: realtime,
		Date:            time.Now(),
		Avg:             &t.Temperature,
	}

	return result
}

type RealtimeHumidityEntity struct {
	Humidity float64 `json:"humidity"`
}

func (e RealtimeHumidityEntity) ConvToModel() (result models.Humidity) {
	result = models.Humidity{
		UpdateFrequency: realtime,
		Date:            time.Now(),
		Avg:             &e.Humidity,
	}
	return
}

type RealtimeCloudrateEntity struct {
	Cloudrate float64 `json:"cloudrate"`
}

func (t RealtimeCloudrateEntity) ConvToModel() (result models.CloudRate) {

	result = models.CloudRate{
		UpdateFrequency: realtime,
		Date:            time.Now(),
		Avg:             &t.Cloudrate,
	}

	return result
}

type RealtimeSkyconEntity struct {
	Skycon string `json:"skycon"`
}

func (t RealtimeSkyconEntity) ConvToModel() (result models.Skycon) {

	result = models.Skycon{
		UpdateFrequency: realtime,
		Date:            time.Now(),
		Value:           &t.Skycon,
	}

	return result
}

type RealtimeVisibilityEntity struct {
	Visibility float64 `json:"visibility"`
}

func (t RealtimeVisibilityEntity) ConvToModel() (result models.Visibility) {

	result = models.Visibility{
		UpdateFrequency: realtime,
		Date:            time.Now(),
		Avg:             &t.Visibility,
	}

	return result
}

type RealtimeDswrfEntity struct {
	Dswrf float64 `json:"dswrf"`
}

func (t RealtimeDswrfEntity) ConvToModel() (result models.DSWRF) {

	result = models.DSWRF{
		UpdateFrequency: realtime,
		Date:            time.Now(),
		Avg:             &t.Dswrf,
	}

	return result
}

type RealtimeWindEntity struct {
	Wind struct {
		Speed     float64 `json:"speed"`
		Direction float64 `json:"direction"`
	} `json:"wind"`
}

func (t RealtimeWindEntity) ConvToModel() (result models.Wind) {

	result = models.Wind{
		UpdateFrequency: realtime,
		Date:            time.Now(),
		AvgSpeed:        &t.Wind.Speed,
		AvgDirection:    &t.Wind.Direction,
	}

	return result
}

type RealtimePressureEntity struct {
	Pressure float64 `json:"pressure"`
}

func (t RealtimePressureEntity) ConvToModel() (result models.Pressure) {

	result = models.Pressure{
		UpdateFrequency: realtime,
		Date:            time.Now(),
		Avg:             &t.Pressure,
	}

	return result
}

type RealtimeApparentTemperatureEntity struct {
	ApparentTemperature float64 `json:"apparent_temperature"`
}

func (t RealtimeApparentTemperatureEntity) ConvToModel() (result models.ApparentTemperature) {

	result = models.ApparentTemperature{
		UpdateFrequency: realtime,
		Date:            time.Now(),
		Avg:             &t.ApparentTemperature,
	}

	return result
}

type RealtimeUltravioletEntity struct {
	Ultraviolet struct {
		Index float64 `json:"index"`
		Desc  string  `json:"desc"`
	} `json:"ultraviolet"`
}

func (t RealtimeUltravioletEntity) ConvToModel() (result models.Ultraviolet) {

	sprint := fmt.Sprint(t.Ultraviolet.Index)
	result = models.Ultraviolet{
		UpdateFrequency: realtime,
		Date:            time.Now(),
		Index:           &sprint,
		Description:     &t.Ultraviolet.Desc,
	}

	return result
}

type RealtimeComfortEntity struct {
	Comfort struct {
		Index float64 `json:"index"`
		Desc  string  `json:"desc"`
	} `json:"comfort"`
}

func (t RealtimeComfortEntity) ConvToModel() (result models.Comfort) {

	sprint := fmt.Sprint(t.Comfort.Index)
	result = models.Comfort{
		UpdateFrequency: realtime,
		Date:            time.Now(),
		Description:     &t.Comfort.Desc,
		Index:           &sprint,
	}

	return result
}

type RealtimePm25Entity struct {
	Pm25 float64 `json:"pm25"`
}

func (t RealtimePm25Entity) ConvToModel() (result models.PM25) {

	result = models.PM25{
		UpdateFrequency: realtime,
		Date:            time.Now(),
		Avg:             &t.Pm25,
	}

	return result
}

type RealtimePm10Entity struct {
	Pm10 float64 `json:"pm10"`
}

func (t RealtimePm10Entity) ConvToModel() (result models.PM10) {

	result = models.PM10{
		Date: time.Now(),
		Avg:  &t.Pm10,
	}

	return result
}

type RealtimeO3Entity struct {
	O3 float64 `json:"o3"`
}

func (t RealtimeO3Entity) ConvToModel() (result models.O3) {

	result = models.O3{
		Date: time.Now(),
		Avg:  &t.O3,
	}

	return result
}

type RealtimeSo2Entity struct {
	So2 float64 `json:"so2"`
}

func (t RealtimeSo2Entity) ConvToModel() (result models.So2) {

	result = models.So2{
		Date: time.Now(),
		Avg:  &t.So2,
	}

	return result
}

type RealtimeNo2Entity struct {
	No2 float64 `json:"no2"`
}

func (t RealtimeNo2Entity) ConvToModel() (result models.No2) {

	result = models.No2{
		Date: time.Now(),
		Avg:  &t.No2,
	}

	return result
}

type RealtimeCoEntity struct {
	Co float64 `json:"co"`
}

func (t RealtimeCoEntity) ConvToModel() (result models.Co) {

	result = models.Co{
		Date: time.Now(),
		Avg:  &t.Co,
	}

	return result
}

type RealtimeAqiEntity struct {
	Aqi struct {
		Chn float64 `json:"chn"`
		Usa float64 `json:"usa"`
	} `json:"aqi"`
}

func (t RealtimeAqiEntity) ConvToModel() (result models.AQI) {

	result = models.AQI{
		UpdateFrequency: realtime,
		Date:            time.Now(),
		AvgCHN:          &t.Aqi.Chn,
		AvgUSA:          &t.Aqi.Usa,
	}

	return result
}

type RealTimeEntity struct {
	Realtime struct {
		Status string `json:"status"`
		RealtimeTemperatureEntity
		RealtimeHumidityEntity
		RealtimeCloudrateEntity
		RealtimeSkyconEntity
		RealtimeVisibilityEntity
		RealtimeDswrfEntity
		RealtimeWindEntity
		RealtimePressureEntity
		RealtimeApparentTemperatureEntity
		RealtimePrecipitationEntity
		AirQuality struct {
			RealtimePm25Entity
			RealtimePm10Entity
			RealtimeO3Entity
			RealtimeSo2Entity
			RealtimeNo2Entity
			RealtimeCoEntity
			RealtimeAqiEntity
			Description struct {
				Chn string `json:"chn"`
				Usa string `json:"usa"`
			} `json:"description"`
		} `json:"air_quality"`
		LifeIndex struct {
			RealtimeUltravioletEntity
			RealtimeComfortEntity
		} `json:"life_index"`
	} `json:"realtime"`
}

func (r *RealTimeEntity) Save() {
	temperature := r.Realtime.RealtimeTemperatureEntity.ConvToModel()
	humidity := r.Realtime.RealtimeHumidityEntity.ConvToModel()
	cloudRate := r.Realtime.RealtimeCloudrateEntity.ConvToModel()
	skycon := r.Realtime.RealtimeSkyconEntity.ConvToModel()
	visibility := r.Realtime.RealtimeVisibilityEntity.ConvToModel()
	dswrf := r.Realtime.RealtimeDswrfEntity.ConvToModel()
	wind := r.Realtime.RealtimeWindEntity.ConvToModel()
	pressure := r.Realtime.RealtimePressureEntity.ConvToModel()
	apparentTemperature := r.Realtime.RealtimeApparentTemperatureEntity.ConvToModel()
	precipitation, realTimePrecipitation := r.Realtime.RealtimePrecipitationEntity.ConvToModel()
	pm25 := r.Realtime.AirQuality.RealtimePm25Entity.ConvToModel()
	pm10 := r.Realtime.AirQuality.RealtimePm10Entity.ConvToModel()
	o3 := r.Realtime.AirQuality.RealtimeO3Entity.ConvToModel()
	so2 := r.Realtime.AirQuality.RealtimeSo2Entity.ConvToModel()
	no2 := r.Realtime.AirQuality.RealtimeNo2Entity.ConvToModel()
	co := r.Realtime.AirQuality.RealtimeCoEntity.ConvToModel()
	aqi := r.Realtime.AirQuality.RealtimeAqiEntity.ConvToModel()

	dao.DB.Create(&temperature)
	dao.DB.Create(&humidity)
	dao.DB.Create(&cloudRate)
	dao.DB.Create(&skycon)
	dao.DB.Create(&visibility)
	dao.DB.Create(&dswrf)
	dao.DB.Create(&wind)
	dao.DB.Create(&pressure)
	dao.DB.Create(&apparentTemperature)
	dao.DB.Create(&precipitation)
	dao.DB.Create(&realTimePrecipitation)
	dao.DB.Create(&pm25)
	dao.DB.Create(&pm10)
	dao.DB.Create(&o3)
	dao.DB.Create(&so2)
	dao.DB.Create(&no2)
	dao.DB.Create(&co)
	dao.DB.Create(&aqi)
}
