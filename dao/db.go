package dao

import (
	"caiyu-go/dao/models"
	"caiyu-go/util"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB = &gorm.DB{}
)

func init() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	mysqlFields := util.Fields("Mysql")
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlFields.Get("User"), mysqlFields.Get("Password"),
		mysqlFields.Get("Host"), mysqlFields.Get("Port"), mysqlFields.Get("Database"))

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB = db

}

func AutoMigrateTables() error {
	autoMigrateTables := []interface{}{
		&models.AQI{},
		&models.WeatherAlert{},
		&models.DailyAstro{},
		&models.DailyCarWashing{},
		&models.CloudRate{},
		&models.DailyColdRisk{},
		&models.Comfort{},
		&models.DailyDressing{},
		&models.DSWRF{},
		&models.Humidity{},
		&models.PM25{},
		&models.O3{},
		&models.So2{},
		&models.No2{},
		&models.Co{},
		&models.PM10{},
		&models.DailyPrecipitation08h20h{},
		&models.DailyPrecipitation20h32h{},
		&models.Precipitation{},
		&models.RealTimePrecipitation{},
		&models.Pressure{},
		&models.Skycon{},
		&models.DailySkycon08H20H{},
		&models.DailySkycon20H32H{},
		&models.Temperature{},
		&models.ApparentTemperature{},
		&models.DailyTemperature08H20H{},
		&models.DailyTemperature20H32H{},
		&models.Ultraviolet{},
		&models.Visibility{},
		&models.Wind{},
		&models.DailyWind08h20h{},
		&models.DailyWind20h32h{},
		&models.MonitoringLocation{},
	}

	err := DB.AutoMigrate(autoMigrateTables...)
	if err != nil {
		panic(err)
	}
	return nil
}
