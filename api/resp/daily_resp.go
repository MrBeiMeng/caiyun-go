package resp

import "caiyu-go/api/entity"

type DailyResp struct {
	entity.RespEntity
	Result struct {
		entity.AlertEntity
		entity.DailyEntity
		Primary int `json:"primary"`
	} `json:"result"`
}

func (d *DailyResp) Save() {
	d.Result.AlertEntity.Save()
	d.Result.DailyEntity.Save()
}
