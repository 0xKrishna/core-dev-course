package main

import (
	"reflect"
	"testing"
)

func TestPic(t *testing.T) {
	type args struct {
		dx int
		dy int
	}

	tests := []struct {
		name string
		args args
		want [][]uint8
	}{
		{"empty", args{0, 0}, [][]uint8{}},
		{"one", args{1, 1}, [][]uint8{{0}}},
		{"two", args{2, 2}, [][]uint8{{0, 1}, {1, 0}}},
		{"three", args{3, 3}, [][]uint8{{0, 1, 2}, {1, 0, 3}, {2, 3, 0}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Pic(tt.args.dx, tt.args.dy); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRunMain(t *testing.T) {
	main()
}
