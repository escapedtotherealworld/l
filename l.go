package l

import "reflect"

func reflectList(list interface{}) reflect.Value {
	reflected := reflect.ValueOf(list)
	kind := reflected.Kind()
	switch kind {
	case reflect.Array, reflect.Slice:
		return reflected
	default:
		panic("Invalid data kind")
	}
}

// Exclude every occurency of items in list
func Exclude(list interface{}, items ...interface{}) interface{} {
	reflected := reflectList(list)
	resultList := []interface{}{}
	for i := 0; i < reflected.Len(); i++ {
		c := reflected.Index(i).Interface()
		if InList(items, c) {
			continue
		}
		resultList = append(resultList, c)
	}
	return resultList
}

// InList returns presens state of item in list
func InList(list interface{}, item interface{}) bool {
	reflected := reflectList(list)
	for i := 0; i < reflected.Len(); i++ {
		if reflected.Index(i).Interface() == item {
			return true
		}
	}
	return false
}
