package api

import (
	"caiyu-go/api/entity"
	"caiyu-go/api/resp"
	"caiyu-go/logging"
	"fmt"
)

func GetHourly(location []float64, hourlySteps int) entity.HourlyEntity {
	resp := resp.HourlyResp{}

	params := make(map[string]string)
	params["hourlysteps"] = fmt.Sprintf("%d", hourlySteps)

	CaiYunGet("hourly", location, &resp, params)

	if resp.Status != "ok" {
		logging.Fatal("resp.Status != \"ok\"", resp)
		return entity.HourlyEntity{}
	}

	return resp.Result.HourlyEntity
}
