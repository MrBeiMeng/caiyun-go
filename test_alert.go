package main

import (
	. "caiyu-go/core/alert_util"
	"fmt"
	"math/rand"
)

type Person struct {
	Min int
	Max int
}

func main() {

	personList := make([]Person, 0)

	for i := 0; i < 100; i++ {
		personList = append(personList, Person{Min: rand.Intn(100), Max: rand.Intn(200)})
	}

	NewAlertWrapper(personList).Inject(
		(AVal("Max").Sub(AVal("Min"))).BiggerThan(12).Then(func(entity2 interface{}) interface{} {
			entity := entity2.(Person)
			fmt.Printf("查询到max-min> 12 的person =>max=%d,min=%d,sub=%d \n", entity.Max, entity.Min, entity.Max-entity.Min)
			return nil
		}),
	)
	//AVal[Person]("Min").SmallerThan(9).Then(func(entity Person) interface{} {
	//
	//	fmt.Printf("查询到min<9 的person =>max=%d,min=%d,sub=%d \n", entity.Max, entity.Min, entity.Max-entity.Min)
	//	return nil
	//}),
}
