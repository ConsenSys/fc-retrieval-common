package slice

import (
	"fmt"
	"reflect"
)

func Exists(slice interface{}, item interface{}) (bool, int, error) {
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return false, -1, fmt.Errorf("method 'exists' is designed for a slice and can not operate on %v", v.Kind().String())
	}
	for i := 0; i < v.Len(); i++ {
		if v.Index(i).Interface() == item {
			return true, i, nil
		}
	}
	return false, -1, nil
}
