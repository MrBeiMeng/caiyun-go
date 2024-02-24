package api

import (
	"caiyu-go/api/entity"
	"caiyu-go/api/resp"
	"caiyu-go/logging"
	"fmt"
)

func GetDaily(location []float64, dailySteps int) entity.DailyEntity {
	resp := resp.DailyResp{}

	params := make(map[string]string)
	params["dailysteps"] = fmt.Sprintf("%d", dailySteps)

	CaiYunGet("daily", location, &resp, params)

	if resp.Status != "ok" {
		logging.Fatal("resp.Status != \"ok\"", resp)
		return entity.DailyEntity{}
	}

	return resp.Result.DailyEntity
}
