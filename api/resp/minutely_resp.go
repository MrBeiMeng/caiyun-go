package resp

import "caiyu-go/api/entity"

type MinutelyResp struct {
	entity.RespEntity
	Result struct {
		entity.AlertEntity
		entity.MinutelyEntity
		Primary          int    `json:"primary"`
		ForecastKeypoint string `json:"forecast_keypoint"`
	} `json:"result"`
}

func (h *MinutelyResp) Save() {
	h.Result.AlertEntity.Save()
	h.Result.MinutelyEntity.Save()
}
