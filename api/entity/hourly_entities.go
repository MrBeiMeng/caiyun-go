package entity

import (
	"caiyu-go/dao"
	"caiyu-go/dao/models"
	"caiyu-go/system_common"
	"time"
)

type HourlyPrecipitationEntity struct {
	Precipitation []struct {
		Datetime    string  `json:"datetime"`
		Value       float64 `json:"value"`
		Probability float64 `json:"probability"`
	} `json:"precipitation"`
}

func (d *HourlyPrecipitationEntity) ConvToModel() (result []models.Precipitation) {
	for i := range d.Precipitation {
		element := d.Precipitation[i]

		elementTime, _ := time.Parse(system_common.Layout, element.Datetime)
		result = append(result, models.Precipitation{
			UpdateFrequency: hourly,
			Date:            elementTime,
			Avg:             &element.Value,
			Probability:     &element.Probability,
		})
	}

	return result
}

type HourlyTemperatureEntity struct {
	Temperature []struct {
		Datetime string  `json:"datetime"`
		Value    float64 `json:"value"`
	} `json:"temperature"`
}

func (t HourlyTemperatureEntity) ConvToModel() (result []models.Temperature) {
	for i := range t.Temperature {
		element := t.Temperature[i]

		elementTime, _ := time.Parse(system_common.Layout, element.Datetime)
		result = append(result, models.Temperature{
			UpdateFrequency: hourly,
			Date:            elementTime,
			Avg:             &element.Value,
		})
	}

	return result
}

type HourlyApparentTemperatureEntity struct {
	ApparentTemperature []struct {
		Datetime string  `json:"datetime"`
		Value    float64 `json:"value"`
	} `json:"apparent_temperature"`
}

func (t HourlyApparentTemperatureEntity) ConvToModel() (result []models.ApparentTemperature) {
	for i := range t.ApparentTemperature {
		element := t.ApparentTemperature[i]

		elementTime, _ := time.Parse(system_common.Layout, element.Datetime)
		result = append(result, models.ApparentTemperature{
			UpdateFrequency: hourly,
			Date:            elementTime,
			Avg:             &element.Value,
		})
	}

	return result
}

type HourlyWindEntity struct {
	Wind []struct {
		Datetime  string  `json:"datetime"`
		Speed     float64 `json:"speed"`
		Direction float64 `json:"direction"`
	} `json:"wind"`
}

func (w HourlyWindEntity) ConvToModel() (result []models.Wind) {
	for i := range w.Wind {
		element := w.Wind[i]
		elementTime, _ := time.Parse(system_common.Layout, element.Datetime)
		result = append(result, models.Wind{
			UpdateFrequency: hourly,
			Date:            elementTime,
			AvgSpeed:        &element.Speed,
			AvgDirection:    &element.Direction,
		})
	}

	return result
}

type HourlyHumidityEntity struct {
	Humidity []struct {
		Datetime string  `json:"datetime"`
		Value    float64 `json:"value"`
	} `json:"humidity"`
}

func (d HourlyHumidityEntity) ConvToModel() (result []models.Humidity) {
	for i := range d.Humidity {
		element := d.Humidity[i]
		elementTime, _ := time.Parse(system_common.Layout, element.Datetime)
		result = append(result, models.Humidity{
			UpdateFrequency: hourly,
			Date:            elementTime,
			Avg:             &element.Value,
		})

	}

	return result
}

type HourlyCloudrateEntity struct {
	Cloudrate []struct {
		Datetime string  `json:"datetime"`
		Value    float64 `json:"value"`
	} `json:"cloudrate"`
}

func (d HourlyCloudrateEntity) ConvToModel() (result []models.CloudRate) {
	for i := range d.Cloudrate {
		element := d.Cloudrate[i]
		elementTime, _ := time.Parse(system_common.Layout, element.Datetime)
		result = append(result, models.CloudRate{
			UpdateFrequency: hourly,
			Date:            elementTime,
			Avg:             &element.Value,
		})

	}

	return result
}

type HourlySkyconEntity struct {
	Skycon []struct {
		Datetime string `json:"datetime"`
		Value    string `json:"value"`
	} `json:"skycon"`
}

func (d HourlySkyconEntity) ConvToModel() (result []models.Skycon) {
	for i := range d.Skycon {
		element := d.Skycon[i]
		elementTime, _ := time.Parse(system_common.Layout, element.Datetime)
		result = append(result, models.Skycon{
			UpdateFrequency: hourly,
			Date:            elementTime,
			Value:           &element.Value,
		})

	}

	return result
}

type HourlyPressureEntity struct {
	Pressure []struct {
		Datetime string  `json:"datetime"`
		Value    float64 `json:"value"`
	} `json:"pressure"`
}

func (d HourlyPressureEntity) ConvToModel() (result []models.Pressure) {
	for i := range d.Pressure {
		element := d.Pressure[i]
		elementTime, _ := time.Parse(system_common.Layout, element.Datetime)
		result = append(result, models.Pressure{
			UpdateFrequency: hourly,
			Date:            elementTime,
			Avg:             &element.Value,
		})

	}

	return result
}

type HourlyVisibilityEntity struct {
	Visibility []struct {
		Datetime string  `json:"datetime"`
		Value    float64 `json:"value"`
	} `json:"visibility"`
}

func (d HourlyVisibilityEntity) ConvToModel() (result []models.Visibility) {
	for i := range d.Visibility {
		element := d.Visibility[i]
		elementTime, _ := time.Parse(system_common.Layout, element.Datetime)
		result = append(result, models.Visibility{
			UpdateFrequency: hourly,
			Date:            elementTime,
			Avg:             &element.Value,
		})

	}

	return result
}

type HourlyDswrfEntity struct {
	Dswrf []struct {
		Datetime string  `json:"datetime"`
		Value    float64 `json:"value"`
	} `json:"dswrf"`
}

func (d HourlyDswrfEntity) ConvToModel() (result []models.DSWRF) {
	for i := range d.Dswrf {
		element := d.Dswrf[i]
		elementTime, _ := time.Parse(system_common.Layout, element.Datetime)
		result = append(result, models.DSWRF{
			UpdateFrequency: hourly,
			Date:            elementTime,
			Avg:             &element.Value,
		})

	}

	return result
}

type HourlyAqiEntity struct {
	Aqi []struct {
		Datetime string `json:"datetime"`
		Value    struct {
			Chn float64 `json:"chn"`
			Usa float64 `json:"usa"`
		} `json:"value"`
	} `json:"aqi"`
}

func (d HourlyAqiEntity) ConvToModel() (result []models.AQI) {
	for i := range d.Aqi {
		element := d.Aqi[i]
		elementTime, _ := time.Parse(system_common.Layout, element.Datetime)
		result = append(result, models.AQI{
			UpdateFrequency: hourly,
			Date:            elementTime,
			AvgCHN:          &element.Value.Chn,
			AvgUSA:          &element.Value.Usa,
		})

	}

	return result
}

type HourlyPm25Entity struct {
	Pm25 []struct {
		Datetime string  `json:"datetime"`
		Value    float64 `json:"value"`
	} `json:"pm25"`
}

func (d HourlyPm25Entity) ConvToModel() (result []models.PM25) {
	for i := range d.Pm25 {
		element := d.Pm25[i]
		elementTime, _ := time.Parse(system_common.Layout, element.Datetime)
		result = append(result, models.PM25{
			UpdateFrequency: hourly,
			Date:            elementTime,
			Avg:             &element.Value,
		})

	}

	return result
}

type HourlyEntity struct {
	Hourly struct {
		Status      string `json:"status"`
		Description string `json:"description"`
		HourlyPrecipitationEntity
		HourlyTemperatureEntity
		HourlyApparentTemperatureEntity
		HourlyWindEntity
		HourlyHumidityEntity
		HourlyCloudrateEntity
		HourlySkyconEntity
		HourlyPressureEntity
		HourlyVisibilityEntity
		HourlyDswrfEntity
		AirQuality struct {
			HourlyAqiEntity
			HourlyPm25Entity
		} `json:"air_quality"`
	} `json:"hourly"`
}

func (h *HourlyEntity) Save() {
	precipitations := h.Hourly.HourlyPrecipitationEntity.ConvToModel()
	temperatures := h.Hourly.HourlyTemperatureEntity.ConvToModel()
	apparentTemperatures := h.Hourly.HourlyApparentTemperatureEntity.ConvToModel()
	winds := h.Hourly.HourlyWindEntity.ConvToModel()
	humidities := h.Hourly.HourlyHumidityEntity.ConvToModel()
	cloudRates := h.Hourly.HourlyCloudrateEntity.ConvToModel()
	skycons := h.Hourly.HourlySkyconEntity.ConvToModel()
	pressures := h.Hourly.HourlyPressureEntity.ConvToModel()
	visibilities := h.Hourly.HourlyVisibilityEntity.ConvToModel()
	dswrfs := h.Hourly.HourlyDswrfEntity.ConvToModel()
	aqis := h.Hourly.AirQuality.HourlyAqiEntity.ConvToModel()
	pm25s := h.Hourly.AirQuality.HourlyPm25Entity.ConvToModel()

	dao.DB.Create(&precipitations)
	dao.DB.Create(&temperatures)
	dao.DB.Create(&apparentTemperatures)
	dao.DB.Create(&winds)
	dao.DB.Create(&humidities)
	dao.DB.Create(&cloudRates)
	dao.DB.Create(&skycons)
	dao.DB.Create(&pressures)
	dao.DB.Create(&visibilities)
	dao.DB.Create(&dswrfs)
	dao.DB.Create(&aqis)
	dao.DB.Create(&pm25s)
}
