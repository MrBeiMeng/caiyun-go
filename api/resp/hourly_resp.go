package resp

import "caiyu-go/api/entity"

type HourlyResp struct {
	entity.RespEntity
	Result struct {
		entity.AlertEntity
		entity.HourlyEntity
		Primary          int    `json:"primary"`
		ForecastKeypoint string `json:"forecast_keypoint"`
	} `json:"result"`
}

func (h *HourlyResp) Save() {
	h.Result.AlertEntity.Save()
	h.Result.HourlyEntity.Save()
}
