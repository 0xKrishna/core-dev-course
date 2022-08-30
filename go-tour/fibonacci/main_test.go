package main

import (
	"reflect"
	"testing"
)

func Test_fibonacci(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"fibonacci(1)", 1},
		{"fibonacci(2)", 1},
		{"fibonacci(3)", 2},
		{"fibonacci(4)", 3},
		{"fibonacci(5)", 5},
		{"fibonacci(6)", 8},
		{"fibonacci(7)", 13},
		{"fibonacci(8)", 21},
		{"fibonacci(9)", 34},
		{"fibonacci(10)", 55},
	}

	f := fibonacci()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRunMain(t *testing.T) {
	main()
}
