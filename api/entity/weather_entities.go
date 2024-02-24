package entity

type WeatherEntity struct {
	AlertEntity
	RealTimeEntity
	MinutelyEntity
	HourlyEntity
	DailyEntity
}

func (w *WeatherEntity) Save() {
	w.AlertEntity.Save()
	w.RealTimeEntity.Save()
	w.MinutelyEntity.Save()
	w.HourlyEntity.Save()
	w.DailyEntity.Save()
}
