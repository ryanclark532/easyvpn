package utils

import (
	"fmt"
	"reflect"
)

func CheckEmpty[T interface{}](emptyStruct T, varStruct T) bool {
	fmt.Println(emptyStruct)
	fmt.Println(varStruct)
	fmt.Println(reflect.DeepEqual(emptyStruct, varStruct))
	return reflect.DeepEqual(emptyStruct, varStruct)
}
