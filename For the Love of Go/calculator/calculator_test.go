package calculator_test

import (
	"calculator"
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 1, b: 1, want: 2},
		{a: 5, b: 0, want: 5},
	}

	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Add(%f,%f): want %f,got %f", tc.a, tc.b, tc.want, got)
		}
	}

}

func TestSubtract(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 2, b: 2, want: 0},
		{a: 1, b: 2, want: -1},
		{a: 5, b: 0, want: 5},
	}

	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Subtract(%f,%f): want %f,got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 1, b: 1, want: 1},
		{a: 5, b: 0, want: 0},
		{a: -5, b: -2, want: 10},
		{a: -2, b: 3, want: -6},
		{a: 10, b: 0.5, want: 5},
	}

	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Multiply(%f,%f): want %f,got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func closeEnough(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

func TestDivide(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 2, b: 2, want: 1},
		{a: -1, b: -1, want: 1},
		{a: 10, b: 5, want: 2},
		{a: 1, b: 3, want: 0.3333},
	}

	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		if err != nil {
			t.Fatalf("want no error for vaild input, got %v", err)
		}
		if !closeEnough(tc.want, got, 0.001) {
			t.Errorf("Divide(%f,%f): want %f,got %f", tc.a, tc.b, tc.want, got)
		}
	}

}

func TestDivideInvalid(t *testing.T) {
	t.Parallel()
	_, err := calculator.Divide(2, 0)
	if err == nil {
		t.Errorf("we want error for invalid input got nil")
	}
}

func TestSqrt(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a    float64
		want float64
	}
	testCases := []testCase{
		{a: 4, want: 2},
		{a: 2, want: 1.41},
		{a: 16, want: 4},
		{a: 0, want: 0},
	}

	for _, tc := range testCases {
		got, err := calculator.Sqrt(tc.a)
		if err != nil {
			t.Fatalf("want no error for vaild input, got %v", err)
		}
		if !closeEnough(tc.want, got, 0.1) {
			t.Errorf("Sqrt(%f): want %f,got %f", tc.a, tc.want, got)
		}
	}
}

func TestSqrtInvalid(t *testing.T) {
	t.Parallel()
	_, err := calculator.Sqrt(-5)
	if err == nil {
		t.Errorf("we want error for invalid input got nil")
	}
}
