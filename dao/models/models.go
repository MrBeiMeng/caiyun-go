package models

import (
	"gorm.io/gorm"
	"time"
)

// AQI represents the  air quality index.
type AQI struct {
	gorm.Model
	UpdateFrequency string    `gorm:"type:varchar(255);not null;"` // daily;hourly:realtime
	Date            time.Time `gorm:"type:datetime;not null"`
	MaxCHN          *float64  `gorm:"type:float"`
	MaxUSA          *float64  `gorm:"type:float"`
	AvgCHN          *float64  `gorm:"type:float"`
	AvgUSA          *float64  `gorm:"type:float"`
	MinCHN          *float64  `gorm:"type:float"`
	MinUSA          *float64  `gorm:"type:float"`
}

// WeatherAlert 定义了气象预警信息的GORM模型
type WeatherAlert struct {
	gorm.Model              // GORM的Model包含ID, CreatedAt, UpdatedAt等字段
	Province      string    `gorm:"size:100"` // 省份
	City          string    `gorm:"size:100"` // 城市
	County        string    `gorm:"size:100"` // 区县，可能为空
	Status        string    `gorm:"size:50"`  // 预警状态
	Code          string    `gorm:"size:20"`  // 预警代码
	Description   string    `gorm:"size:500"` // 预警描述
	PubTimestamp  time.Time // 发布时间戳
	Latitude      float64   `gorm:"type:decimal(10,7)"` // 纬度
	Longitude     float64   `gorm:"type:decimal(10,7)"` // 经度
	AlertID       string    `gorm:"size:100"`           // 预警ID
	Title         string    `gorm:"size:200"`           // 预警标题
	Adcode        string    `gorm:"size:20"`            // 行政区划代码
	Source        string    `gorm:"size:200"`           // 预警信息来源
	Location      string    `gorm:"size:200"`           // 详细位置
	RequestStatus string    `gorm:"size:50"`            // 请求状态
}

// DailyAstro Astro represents the  astronomical information. 日出日落时间
type DailyAstro struct {
	gorm.Model
	Date        time.Time  `gorm:"type:datetime;not null"`
	SunriseTime *time.Time `gorm:"type:time"`
	SunsetTime  *time.Time `gorm:"type:time"`
}

// DailyCarWashing CarWashing represents the  car washing index.
type DailyCarWashing struct {
	gorm.Model
	Date        time.Time `gorm:"type:datetime;not null"`
	Index       *string   `gorm:"type:varchar(10)"`
	Description *string   `gorm:"type:text"`
}

// CloudRate  represents the  cloud coverage information.
type CloudRate struct {
	gorm.Model
	UpdateFrequency string    `gorm:"type:varchar(255);not null;"` // daily;hourly:realtime
	Date            time.Time `gorm:"type:datetime;not null"`
	Max             *float64  `gorm:"type:float"`
	Min             *float64  `gorm:"type:float"`
	Avg             *float64  `gorm:"type:float"`
}

// DailyColdRisk ColdRisk represents the  cold risk index.
type DailyColdRisk struct {
	gorm.Model
	Date        time.Time `gorm:"type:datetime;not null"`
	Index       *string   `gorm:"type:varchar(10)"`
	Description *string   `gorm:"type:text"`
}

// Comfort represents the  comfort index.
type Comfort struct {
	gorm.Model
	UpdateFrequency string    `gorm:"type:varchar(255);not null;"` // daily;realtime
	Date            time.Time `gorm:"type:datetime;not null"`
	Index           *string   `gorm:"type:varchar(10)"`
	Description     *string   `gorm:"type:text"`
}

// DailyDressing Dressing represents the  dressing index.
type DailyDressing struct {
	gorm.Model
	Date        time.Time `gorm:"type:datetime;not null"`
	Index       *string   `gorm:"type:varchar(10)"`
	Description *string   `gorm:"type:text"`
}

// DSWRF represents the  downward shortwave radiation flux.向下短波辐射通量(W/M2)
type DSWRF struct {
	gorm.Model
	UpdateFrequency string    `gorm:"type:varchar(255);not null;"` // daily;hourly:realtime
	Date            time.Time `gorm:"type:datetime;not null"`
	Max             *float64  `gorm:"type:float"`
	Min             *float64  `gorm:"type:float"`
	Avg             *float64  `gorm:"type:float"`
}

// Humidity  represents the  humidity data.地表 2 米相对湿度(%)
type Humidity struct {
	gorm.Model
	UpdateFrequency string    `gorm:"type:varchar(255);not null;"` // daily;hourly:realtime
	Date            time.Time `gorm:"type:datetime;not null"`
	Max             *float64  `gorm:"type:float"`
	Min             *float64  `gorm:"type:float"`
	Avg             *float64  `gorm:"type:float"`
}

// PM25 represents the  PM2.5 data.
type PM25 struct {
	gorm.Model
	UpdateFrequency string    `gorm:"type:varchar(255);not null;"` // daily;hourly:realtime
	Date            time.Time `gorm:"type:datetime;not null"`
	Max             *float64  `gorm:"type:float"`
	Avg             *float64  `gorm:"type:float"`
	Min             *float64  `gorm:"type:float"`
}

type PM10 struct {
	gorm.Model
	Date time.Time `gorm:"type:datetime;not null"`
	Avg  *float64  `gorm:"type:float"`
}

type O3 struct {
	gorm.Model
	Date time.Time `gorm:"type:datetime;not null"`
	Avg  *float64  `gorm:"type:float"`
}

type So2 struct {
	gorm.Model
	Date time.Time `gorm:"type:datetime;not null"`
	Avg  *float64  `gorm:"type:float"`
}

type No2 struct {
	gorm.Model
	Date time.Time `gorm:"type:datetime;not null"`
	Avg  *float64  `gorm:"type:float"`
}

type Co struct {
	gorm.Model
	Date time.Time `gorm:"type:datetime;not null"`
	Avg  *float64  `gorm:"type:float"`
}

// Precipitation represents the  precipitation data from 0h to 24h (today). // 降水量以及降水概率
type Precipitation struct {
	gorm.Model
	UpdateFrequency string    `gorm:"type:varchar(255);not null;"` // daily;hourly:realtime
	Date            time.Time `gorm:"type:datetime;not null"`
	Max             *float64  `gorm:"type:float"`
	Min             *float64  `gorm:"type:float"`
	Avg             *float64  `gorm:"type:float"`
	Probability     *float64  `gorm:"type:float"`
}

// RealTimePrecipitation  represents the  precipitation data from 0h to 24h (today). // 降水量以及降水概率
type RealTimePrecipitation struct {
	gorm.Model
	Date             time.Time `gorm:"type:datetime;not null"`
	LocalIntensity   *float64  `gorm:"type:float"`         // 当地降水强度
	LocalDataSource  *string   `gorm:"type:varchar(255);"` // 数据源
	NearestDistance  *float64  `gorm:"type:float"`         // 最近的降水距离
	NearestIntensity *float64  `gorm:"type:float"`         // 降水强度
}

// DailyPrecipitation08h20h Precipitation08h20h represents the  precipitation data from 08h to 20h.
type DailyPrecipitation08h20h struct {
	gorm.Model
	Date        time.Time `gorm:"type:datetime;not null"`
	Max         *float64  `gorm:"type:float"`
	Min         *float64  `gorm:"type:float"`
	Avg         *float64  `gorm:"type:float"`
	Probability *float64  `gorm:"type:float"`
}

// DailyPrecipitation20h32h Precipitation20h32h represents the  precipitation data from 20h to 32h (next day).
type DailyPrecipitation20h32h struct {
	gorm.Model
	Date        time.Time `gorm:"type:datetime;not null"`
	Max         *float64  `gorm:"type:float"`
	Min         *float64  `gorm:"type:float"`
	Avg         *float64  `gorm:"type:float"`
	Probability *float64  `gorm:"type:float"`
}

// Pressure represents the  atmospheric pressure data.
type Pressure struct {
	gorm.Model
	UpdateFrequency string    `gorm:"type:varchar(255);not null;"` // daily;hourly:realtime
	Date            time.Time `gorm:"type:datetime;not null"`
	Max             *float64  `gorm:"type:float"`
	Min             *float64  `gorm:"type:float"`
	Avg             *float64  `gorm:"type:float"`
}

// Skycon represents the  sky condition data. 云量
type Skycon struct {
	gorm.Model
	UpdateFrequency string    `gorm:"type:varchar(255);not null;"` // daily;hourly:realtime
	Date            time.Time `gorm:"type:datetime;not null"`
	Value           *string   `gorm:"type:varchar(50)"`
}

type DailySkycon08H20H struct {
	gorm.Model
	Date  time.Time `gorm:"type:datetime;not null"`
	Value *string   `gorm:"type:varchar(50)"`
}

type DailySkycon20H32H struct {
	gorm.Model
	Date  time.Time `gorm:"type:datetime;not null"`
	Value *string   `gorm:"type:varchar(50)"`
}

// Temperature represents the  temperature data.
type Temperature struct {
	gorm.Model
	UpdateFrequency string    `gorm:"type:varchar(255);not null;"` // daily;hourly:realtime
	Date            time.Time `gorm:"type:datetime;not null"`
	Max             *float64  `gorm:"type:float;comment:全天地表 2 米气温"`
	Min             *float64  `gorm:"type:float;comment:全天地表 2 米气温"`
	Avg             *float64  `gorm:"type:float;comment:全天地表 2 米气温"`
}

type ApparentTemperature struct {
	gorm.Model
	UpdateFrequency string    `gorm:"type:varchar(255);not null;"` // hourly:realtime
	Date            time.Time `gorm:"type:datetime;not null"`
	Avg             *float64  `gorm:"type:float;comment:全天地表 2 米气温"`
}

type DailyTemperature08H20H struct {
	gorm.Model
	Date time.Time `gorm:"type:datetime;not null"`
	Max  *float64  `gorm:"type:float;comment:白天地表 2 米气温"`
	Min  *float64  `gorm:"type:float;comment:白天地表 2 米气温"`
	Avg  *float64  `gorm:"type:float;comment:白天地表 2 米气温"`
}

type DailyTemperature20H32H struct {
	gorm.Model
	Date time.Time `gorm:"type:datetime;not null"`
	Max  *float64  `gorm:"type:float;comment:夜间地表两米气温"`
	Min  *float64  `gorm:"type:float;comment:夜间地表两米气温"`
	Avg  *float64  `gorm:"type:float;comment:夜间地表两米气温"`
}

// Ultraviolet Ultraviolet represents the  ultraviolet index data.
type Ultraviolet struct {
	gorm.Model
	UpdateFrequency string    `gorm:"type:varchar(255);not null;"` // daily;realtime
	Date            time.Time `gorm:"type:datetime;not null"`
	Index           *string   `gorm:"type:varchar(10)"`
	Description     *string   `gorm:"type:text"`
}

// Visibility represents the  visibility data.
type Visibility struct {
	gorm.Model
	UpdateFrequency string    `gorm:"type:varchar(255);not null;"` // daily;hourly:realtime
	Date            time.Time `gorm:"type:datetime;not null"`
	Max             *float64  `gorm:"type:float"`
	Min             *float64  `gorm:"type:float"`
	Avg             *float64  `gorm:"type:float"`
}

// Wind represents the  wind data.
type Wind struct {
	gorm.Model
	UpdateFrequency string    `gorm:"type:varchar(255);not null;"` // daily;hourly:realtime
	Date            time.Time `gorm:"type:datetime;not null"`
	MaxSpeed        *float64  `gorm:"type:float"`
	MaxDirection    *float64  `gorm:"type:float"`
	MinSpeed        *float64  `gorm:"type:float"`
	MinDirection    *float64  `gorm:"type:float"`
	AvgSpeed        *float64  `gorm:"type:float"`
	AvgDirection    *float64  `gorm:"type:float"`
}

// DailyWind08h20h Wind08h20h represents the  wind data from 08h to 20h.
type DailyWind08h20h struct {
	gorm.Model
	Date         time.Time `gorm:"type:datetime;not null"`
	MaxSpeed     *float64  `gorm:"type:float"`
	MaxDirection *float64  `gorm:"type:float"`
	MinSpeed     *float64  `gorm:"type:float"`
	MinDirection *float64  `gorm:"type:float"`
	AvgSpeed     *float64  `gorm:"type:float"`
	AvgDirection *float64  `gorm:"type:float"`
}

// DailyWind20h32h Wind20h32h represents the  wind data from 20h to 32h (next day).
type DailyWind20h32h struct {
	gorm.Model
	Date         time.Time `gorm:"type:datetime;not null"`
	MaxSpeed     *float64  `gorm:"type:float"`
	MaxDirection *float64  `gorm:"type:float"`
	MinSpeed     *float64  `gorm:"type:float"`
	MinDirection *float64  `gorm:"type:float"`
	AvgSpeed     *float64  `gorm:"type:float"`
	AvgDirection *float64  `gorm:"type:float"`
}

// MonitoringLocation 监控地址
type MonitoringLocation struct {
	gorm.Model
	Latitude  float64 `gorm:"type:decimal(10,7)"` // 纬度
	Longitude float64 `gorm:"type:decimal(10,7)"` // 经度
	Desc      *string `gorm:"type:varchar(255)"`  // 备注
}

type MonitorAction struct {
}
