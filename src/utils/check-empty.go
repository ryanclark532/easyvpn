package utils

import (
	"reflect"
)

func CheckEmpty[T interface{}](emptyStruct T, varStruct T) bool {
	return reflect.DeepEqual(emptyStruct, varStruct)
}
