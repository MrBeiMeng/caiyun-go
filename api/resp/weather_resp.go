package resp

import "caiyu-go/api/entity"

type WeatherResp struct {
	entity.RespEntity
	Result struct {
		entity.AlertEntity
		entity.WeatherEntity
		Primary          int    `json:"primary"`
		ForecastKeypoint string `json:"forecast_keypoint"`
	} `json:"result"`
}

func (w *WeatherResp) Save() {
	w.Result.AlertEntity.Save()
	w.Result.WeatherEntity.Save()
}
