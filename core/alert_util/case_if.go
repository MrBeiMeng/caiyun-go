package alert_util

import (
	"reflect"
)

type AlertCase struct {
	GetVal func(structEntity interface{}) interface{}
}

type AlertVal struct {
	GetVal func(structEntity interface{}) interface{}
}

// AVal 列，将判断传入的属性类型比如和fieldType一致
func AVal(fn string) AlertVal {

	// 默认类型
	a := AlertVal{}

	a.GetVal = func(structEntity interface{}) interface{} {
		valueOf := reflect.ValueOf(structEntity)
		return valueOf.FieldByName(fn).Interface()
	}

	return a
}

func (a AlertVal) And(b AlertVal) AlertVal {
	return AlertVal{
		GetVal: func(structEntity interface{}) interface{} {
			return convToBool(a.GetVal(structEntity)) && convToBool(b.GetVal(structEntity))
		},
	}
}

func (a AlertVal) Sub(b AlertVal) AlertVal {
	return AlertVal{
		GetVal: func(structEntity interface{}) interface{} {
			return convToFloat64(a.GetVal(structEntity)) - convToFloat64(b.GetVal(structEntity))
		},
	}

}

func (a AlertVal) Add(b AlertVal) AlertVal {
	return AlertVal{
		GetVal: func(structEntity interface{}) interface{} {
			return convToFloat64(a.GetVal(structEntity)) + convToFloat64(b.GetVal(structEntity))
		},
	}

}

func (a AlertVal) BiggerThan(val float64) AlertVal {
	return AlertVal{
		GetVal: func(structEntity interface{}) interface{} {
			return convToFloat64(a.GetVal(structEntity)) > val
		},
	}
}

func (a AlertVal) SmallerThan(val float64) AlertVal {
	return AlertVal{
		GetVal: func(structEntity interface{}) interface{} {
			return convToFloat64(a.GetVal(structEntity)) < val
		},
	}
}

func (a AlertVal) Between(x, y float64) AlertVal {
	return AlertVal{
		GetVal: func(structEntity interface{}) interface{} {
			return x <= convToFloat64(a.GetVal(structEntity)) && convToFloat64(a.GetVal(structEntity)) < y
		},
	}
}

func (a AlertVal) Then(do func(entity interface{}) interface{}) AlertCase {

	upperCall := a.GetVal
	a.GetVal = func(structEntity interface{}) interface{} {
		if value, ok := upperCall(structEntity).(bool); !ok || !value {
			if !ok {
				panic("未成立表达式便调用Then方法")
			}
			return nil
		}
		return do(structEntity)
	}

	return AlertCase{
		GetVal: a.GetVal,
	}
}

// (AVal("Max").Sub("Min")).BiggerThan(12).Then(func(entity interface){
// yourEntity := entity.(YourEntity)
// // 你可以构造你的Alert类型
//})
// 可以简写成 表名 Max sub Min biggerThan 12 then '差异值达到了${Max sub Min}'

type AlertWrapper struct {
	data interface{}
}

// NewAlertWrapper 通过数据构造
func NewAlertWrapper(data interface{}) AlertWrapper {
	return AlertWrapper{data: data}
}

// Inject 通过case进行判断
func (a AlertWrapper) Inject(alerts ...AlertCase) {
	// 遍历data，可能是数组也可能是单个，然后在每个元素上执行Case
	dataType := reflect.TypeOf(a.data)

	switch dataType.Kind() {
	case reflect.Struct:
		for _, alert := range alerts {
			alert.GetVal(a.data)
		}

	case reflect.Array, reflect.Slice:
		dataValue := reflect.ValueOf(a.data)
		for i := 0; i < dataValue.Len(); i++ {
			element := dataValue.Index(i).Interface()

			for _, alert := range alerts {
				alert.GetVal(element)
			}
		}
	default:
		panic("error：不可注入的类型")
	}

}
