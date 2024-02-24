package entity

import (
	"caiyu-go/dao"
	"caiyu-go/dao/models"
	"time"
)

type MinutelyEntity struct {
	Minutely struct {
		Status          string    `json:"status"`
		Datasource      string    `json:"datasource"`
		Precipitation2H []float64 `json:"precipitation_2h"`
		Precipitation   []float64 `json:"precipitation"`
		Probability     []float64 `json:"probability"`
		Description     string    `json:"description"`
	} `json:"minutely"`
}

func (d *MinutelyEntity) Save() {
	precipitations := d.ConvToModel()
	dao.DB.Create(&precipitations)
}

func (d *MinutelyEntity) ConvToModel() (result []models.Precipitation) {
	elementTime := time.Now()

	for i := range d.Minutely.Precipitation2H {
		precipitation := d.Minutely.Precipitation2H[i]

		result = append(result, models.Precipitation{
			UpdateFrequency: minutely,
			Date:            elementTime.Add(time.Minute),
			Avg:             &precipitation,
			Probability:     &d.Minutely.Probability[int(i/30)], // 表示每半小时的变化情况
		})
	}

	return result
}
