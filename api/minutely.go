package api

import (
	"caiyu-go/api/entity"
	"caiyu-go/api/resp"
	"caiyu-go/logging"
)

func GetMinutely(location []float64) entity.MinutelyEntity {
	resp := resp.MinutelyResp{}

	CaiYunGet("minutely", location, &resp, nil)

	if resp.Status != "ok" {
		logging.Fatal("resp.Status != \"ok\"", resp)
		return entity.MinutelyEntity{}
	}

	return resp.Result.MinutelyEntity
}
