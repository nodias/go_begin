package reflect

import "reflect"

func NewMap(key, value interface{}) interface{} {
	keyType := reflect.TypeOf(key)
	valueType := reflect.TypeOf(value)
	mapType := reflect.MapOf(keyType, valueType)
	mapValue := reflect.MakeMap(mapType)
	return mapValue.Interface()
}
