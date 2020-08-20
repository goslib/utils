package utlReflect

import "reflect"

func IsInterfaceValueNil(v interface{}) bool {
	return v == nil || reflect.ValueOf(v).IsNil()
}
