package operations

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		x, y, z float64
	}{
		{2, 2, 4},
		{-2, 2, 0},
		{-2, -2, -4},
	}

	for _, tc := range tests {
		got, _ := Add(float64(tc.x), float64(tc.y))
		expected := tc.z

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("Test failed. Performing addition of %v and %v got %v but expected %v", tc.x, tc.y, got, expected)
		}
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		x, y, z float64
	}{
		{2, 2, 0},
		{-2, 2, -4},
		{-2, -2, 0},
	}

	for _, tc := range tests {
		got, _ := Subtract(float64(tc.x), float64(tc.y))
		expected := tc.z

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("Test failed. Performing subtraction of %v and %v got %v but expected %v", tc.x, tc.y, got, expected)
		}
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		x, y, z float64
	}{
		{2, 2, 4},
		{-2, 2, -4},
		{-2, -2, 4},
	}

	for _, tc := range tests {
		got, _ := Multiply(float64(tc.x), float64(tc.y))
		expected := tc.z

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("Test failed. Performing multiplication of %v and %v got %v but expected %v", tc.x, tc.y, got, expected)
		}
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		x, y, z float64
	}{
		{2, 2, 1},
		{-2, 2, -1},
		{-2, -2, 1},
		{2, 0, 0},
	}

	for _, tc := range tests {
		got, _ := Divide(float64(tc.x), float64(tc.y))
		expected := tc.z

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("Test failed. Performing division of %v and %v got %v but expected %v", tc.x, tc.y, got, expected)
		}
	}
}
