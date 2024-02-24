package alert_util2

import (
	"fmt"
	"reflect"
)

func convToFloat64(val interface{}) float64 {
	upperType := reflect.TypeOf(val)

	var floatA float64
	// 转换成float64
	switch upperType.Kind() {
	case reflect.Int:
		floatA = float64(val.(int))
	case reflect.Int8:
		floatA = float64(val.(int8))
	case reflect.Int16:
		floatA = float64(val.(int16))
	case reflect.Int32:
		floatA = float64(val.(int32))
	case reflect.Int64:
		floatA = float64(val.(int64))
	case reflect.Float32:
		floatA = float64(val.(float32))
	case reflect.Float64:
		floatA = val.(float64)
	default:
		fmt.Println(upperType.Kind())
		panic("Unsupported type")
	}

	return floatA
}

func convToBool(val interface{}) bool {
	if _, ok := val.(bool); !ok {
		panic("无法转为布尔类型")
	}

	return val.(bool)
}
