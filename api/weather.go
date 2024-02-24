package api

import (
	"caiyu-go/api/entity"
	"caiyu-go/api/resp"
	"caiyu-go/logging"
	"fmt"
)

func GetWeather(location []float64, dailySteps int, hourlySteps int) entity.WeatherEntity {
	resp := resp.WeatherResp{}

	params := make(map[string]string)
	params["dailysteps"] = fmt.Sprintf("%d", dailySteps)
	params["hourlysteps"] = fmt.Sprintf("%d", hourlySteps)

	CaiYunGet("daily", location, &resp, params)

	if resp.Status != "ok" {
		logging.Fatal("resp.Status != \"ok\"", resp)
		return entity.WeatherEntity{}
	}

	return resp.Result.WeatherEntity
}
