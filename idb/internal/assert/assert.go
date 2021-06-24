// +build js,wasm

package assert

import (
	"reflect"
	"strings"
	"testing"
)

// Error asserts err is not nil
func Error(tb testing.TB, err error) bool {
	tb.Helper()
	if err == nil {
		tb.Error("Expected error, got nil")
		return false
	}
	return true
}

// NoError asserts err is nil
func NoError(tb testing.TB, err error) bool {
	tb.Helper()
	if err != nil {
		tb.Errorf("Unexpected error: %+v", err)
		return false
	}
	return true
}

// Zero asserts value is the zero value
func Zero(tb testing.TB, value interface{}) bool {
	tb.Helper()
	if !reflect.ValueOf(value).IsZero() {
		tb.Errorf("Value should be zero, got: %#v", value)
		return false
	}
	return true
}

// NotZero asserts value is not the zero value
func NotZero(tb testing.TB, value interface{}) bool {
	tb.Helper()
	if reflect.ValueOf(value).IsZero() {
		tb.Error("Value should not be zero")
		return false
	}
	return true
}

// Equal asserts actual is equal to expected
func Equal(tb testing.TB, expected, actual interface{}) bool {
	tb.Helper()
	if !reflect.DeepEqual(expected, actual) {
		tb.Errorf("Expected: %#v\nActual:    %#v", expected, actual)
		return false
	}
	return true
}

// NotEqual asserts actual is not equal to expected
func NotEqual(tb testing.TB, expected, actual interface{}) bool {
	tb.Helper()
	if reflect.DeepEqual(expected, actual) {
		tb.Errorf("Should not be equal.\nExpected: %#v\nActual:    %#v", expected, actual)
		return false
	}
	return true
}

func contains(tb testing.TB, collection, item interface{}) bool {
	collectionVal := reflect.ValueOf(collection)
	switch collectionVal.Kind() {
	case reflect.Slice:
		length := collectionVal.Len()
		for i := 0; i < length; i++ {
			candidateItem := collectionVal.Index(i).Interface()
			if reflect.DeepEqual(candidateItem, item) {
				return true
			}
		}
		return false
	case reflect.String:
		itemVal := reflect.ValueOf(item)
		if itemVal.Kind() != reflect.String {
			tb.Errorf("Invalid item type for string collection. Expected string, got: %T", item)
			return false
		}
		return strings.Contains(collection.(string), item.(string))
	default:
		tb.Errorf("Invalid collection type. Expected slice, got: %T", collection)
		return false
	}
}

// Contains asserts item is contained by collection
func Contains(tb testing.TB, collection, item interface{}) bool {
	tb.Helper()

	if !contains(tb, collection, item) {
		tb.Errorf("Collection does not contain expected item:\nCollection: %#v\nExpected item: %#v", collection, item)
		return false
	}
	return true
}

// NotContains asserts item is not contained by collection
func NotContains(tb testing.TB, collection, item interface{}) bool {
	tb.Helper()

	if contains(tb, collection, item) {
		tb.Errorf("Collection contains unexpected item:\nCollection: %#v\nUnexpected item: %#v", collection, item)
		return false
	}
	return true
}
