package system_common

const (
	Home = 1.1 // 经纬度
)

// ConvL a 经度，b 纬度
func ConvL(location float64) []float64 {
	a := float64(int(location))
	b := location - a

	return []float64{a / 10000, b * 100}
}
