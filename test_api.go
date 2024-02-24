package main

import (
	"caiyu-go/api"
	"caiyu-go/system_common"
)

func testDaily() {
	daily := api.GetDaily(system_common.ConvL(system_common.Home), 15)
	daily.Save()
}

func testHourly() {
	hourly := api.GetHourly(system_common.ConvL(system_common.Home), 360)
	hourly.Save()
}

func testMinutely() {
	minutely := api.GetMinutely(system_common.ConvL(system_common.Home))
	minutely.Save()
}

func testRealtime() {
	realtime := api.GetRealTime(system_common.ConvL(system_common.Home))
	realtime.Save()
}

func testWeather() {
	weather := api.GetWeather(system_common.ConvL(system_common.Home), 15, 360)
	weather.Save()
}

func main() {
	testWeather()
}
