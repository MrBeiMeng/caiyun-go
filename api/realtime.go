package api

import (
	"caiyu-go/api/entity"
	"caiyu-go/api/resp"
	"caiyu-go/logging"
)

func GetRealTime(location []float64) entity.RealTimeEntity {
	realTimeResp := resp.RealTimeResp{}

	CaiYunGet("realtime", location, &realTimeResp, nil)

	if realTimeResp.Status != "ok" {
		logging.Fatal("realTimeResp.Status != \"ok\"", realTimeResp)
		return entity.RealTimeEntity{}
	}

	return realTimeResp.Result.RealTimeEntity
}
