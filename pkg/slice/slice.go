/*
Package slice - implements slice methods, missing in Go's standard library.
E.g. generic search operation.
*/
package slice

import (
	"fmt"
	"reflect"
)

// Exists - generic (applicable for any data type) method, used to check if a slice contains an element.
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
