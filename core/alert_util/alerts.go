package alert_util

const (
	white = iota
	blue
	yellow
	orange
	red // 最高级别预警
)

type WeatherAlert struct {
	//Latitude    float64
	//Longitude   float64
	Description string
	//Level       int
}
