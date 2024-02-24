package resp

import "caiyu-go/api/entity"

type RealTimeResp struct {
	entity.RespEntity
	Result struct {
		entity.AlertEntity
		entity.RealTimeEntity
		Primary int `json:"primary"`
	} `json:"result"`
}

func (r *RealTimeResp) Save() {
	r.Result.AlertEntity.Save()
	r.Result.RealTimeEntity.Save()
}
