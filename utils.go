package utils

import (
	"reflect"
	"testing"
)

func AssertArraysEqual(t *testing.T, expected []int, result []int) {
	if len(expected) != len(result) || !reflect.DeepEqual(expected, result) {
		fail(t, "Arrays were not equal", expected, result)
	}
}

func AssertEqual(t *testing.T, values ...interface{}) {
	if (len(values)) != 2 {
		t.Error("Only expecting two values.")
		return
	}

	expected := values[0]
	result := values[1]
	expectedType := reflect.TypeOf(expected)
	resultType := reflect.TypeOf(result)

	if expectedType.Kind() != resultType.Kind() {
		fail(t, "Items under test are different types.", expectedType, resultType)
		return
	}

	if expected != result {
		fail(t, "Items were not equal", expected, result)
	}
}

func fail(t *testing.T, msg string, expected interface{}, result interface{}) {
	t.Error(msg, "\nExp:", expected, "\nAct:", result)
}
