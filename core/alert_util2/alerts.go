package alert_util2

import "time"

const (
	white = iota
	blue
	yellow
	orange
	red // 最高级别预警
)

type WeatherAlert struct {
	Description string
	Level       int
	Date        time.Time
}

type AWeatherAlert struct {
	Latitude  float64
	Longitude float64
	Data      []WeatherAlert
}

//type Person struct {
//	Max int
//	Min int
//}
