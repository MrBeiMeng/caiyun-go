package entity

type RespEntity struct {
	Status     string    `json:"status"`
	ApiVersion string    `json:"api_version"`
	ApiStatus  string    `json:"api_status"`
	Lang       string    `json:"lang"`
	Unit       string    `json:"unit"`
	Tzshift    float64   `json:"tzshift"`
	Timezone   string    `json:"timezone"`
	ServerTime int       `json:"server_time"`
	Location   []float64 `json:"location"`
}
