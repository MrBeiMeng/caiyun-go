package alert_util2

import (
	"reflect"
)

// AlertVal 结构体属性控制
type AlertVal[T any] struct {
	GetVal func(structEntity T) interface{}
}

func AVal[T any](fn string) AlertVal[T] {
	return AlertVal[T]{
		GetVal: func(structEntity T) interface{} {
			valueOf := reflect.ValueOf(structEntity)
			if valueOf.Kind() != reflect.Struct {
				panic("传入参数非结构体类型")
			}

			fieldValue := valueOf.FieldByName(fn)
			if !fieldValue.IsValid() {
				return nil
			}

			return fieldValue.Interface()
		},
	}
}

func (a AlertVal[T]) And(b AlertVal[T]) AlertVal[T] {
	return AlertVal[T]{
		GetVal: func(structEntity T) interface{} {
			return convToBool(a.GetVal(structEntity)) && convToBool(b.GetVal(structEntity))
		},
	}
}

func (a AlertVal[T]) Add(b AlertVal[T]) AlertVal[T] {
	return AlertVal[T]{
		GetVal: func(structEntity T) interface{} {
			return convToFloat64(a.GetVal(structEntity)) + convToFloat64(b.GetVal(structEntity))
		},
	}
}

func (a AlertVal[T]) Sub(b AlertVal[T]) AlertVal[T] {
	return AlertVal[T]{
		GetVal: func(structEntity T) interface{} {
			return convToFloat64(a.GetVal(structEntity)) - convToFloat64(b.GetVal(structEntity))
		},
	}
}

func (a AlertVal[T]) Between(x, y float64) AlertVal[T] {
	return AlertVal[T]{
		GetVal: func(structEntity T) interface{} {
			return x <= convToFloat64(a.GetVal(structEntity)) && convToFloat64(a.GetVal(structEntity)) < y
		},
	}
}

func (a AlertVal[T]) BiggerThan(val float64) AlertVal[T] {
	return AlertVal[T]{
		GetVal: func(structEntity T) interface{} {
			return convToFloat64(a.GetVal(structEntity)) > val
		},
	}
}

func (a AlertVal[T]) SmallerThan(val float64) AlertVal[T] {
	return AlertVal[T]{
		GetVal: func(structEntity T) interface{} {
			return convToFloat64(a.GetVal(structEntity)) < val
		},
	}
}

func (a AlertVal[T]) Then(do func(entity T) interface{}) AlertCase[T] {

	return AlertCase[T]{
		GetVal: func(structEntity T) interface{} {
			if value, ok := a.GetVal(structEntity).(bool); !ok || !value {
				if !ok {
					panic("Consider returning an error instead of panic for better error handling")
				}
				return nil
			}
			return do(structEntity)
		},
	}
}

// AlertCase 条件表达式
type AlertCase[T any] struct {
	GetVal func(structEntity T) interface{}
}

// AlertWrapper 数据集结构
type AlertWrapper[T any] struct {
	data []T
}

// NewAlertWrapper 数据集
func NewAlertWrapper[T any](data []T) AlertWrapper[T] {
	return AlertWrapper[T]{data: data}
}

// Inject 将条件表达式注入数据集合
func (a AlertWrapper[T]) Inject(alerts ...AlertCase[T]) {
	for _, element := range a.data {
		for _, alert := range alerts {
			alert.GetVal(element)
		}
	}
}
