package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func ConvertIntArrToStrArr(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}

func TestOpcodeOne(t *testing.T) {
	data := []int{1, 0, 0, 0, 99}
	expectedData := []int{2, 0, 0, 0, 99}

	OpcodeOne(0, data)
	if !reflect.DeepEqual(data, expectedData) {
		t.Errorf("expected %s, got %s", ConvertIntArrToStrArr(expectedData, ", "),
			ConvertIntArrToStrArr(data, ", "))
	}
}

func TestOpcodeTWo(t *testing.T) {
	data := []int{2, 3, 0, 3, 99}
	expectedData := []int{2, 3, 0, 6, 99}

	OpcodeTwo(0, data)
	if !reflect.DeepEqual(data, expectedData) {
		t.Errorf("expected %s, got %s", ConvertIntArrToStrArr(expectedData, ", "),
			ConvertIntArrToStrArr(data, ", "))
	}

	data = []int{2, 4, 4, 5, 99, 0}
	expectedData = []int{2, 4, 4, 5, 99, 9801}
	OpcodeTwo(0, data)
	if !reflect.DeepEqual(data, expectedData) {
		t.Errorf("expected %s, got %s", ConvertIntArrToStrArr(expectedData, ", "),
			ConvertIntArrToStrArr(data, ", "))
	}

}

func TestRunProgram(t *testing.T) {
	data := []int{1, 1, 1, 4, 99, 5, 6, 0, 99}
	expectedData := []int{30, 1, 1, 4, 2, 5, 6, 0, 99}
	RunProgram(data)

	if !reflect.DeepEqual(data, expectedData) {
		t.Errorf("expected %s, got %s", ConvertIntArrToStrArr(expectedData, ", "),
			ConvertIntArrToStrArr(data, ", "))
	}

}
