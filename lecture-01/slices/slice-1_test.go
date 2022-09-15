package slices

import (
	"reflect"
	"testing"
)

func TestSlice_1(t *testing.T) {
	type args struct {
		s []int
		N int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{"empty slice", args{[]int{}, 1}, []int{}},
		{"one element", args{[]int{1}, 1}, []int{2}},
		{"two elements", args{[]int{1, 2}, 1}, []int{2, 3}},
		{"three elements", args{[]int{1, 2, 3}, 1}, []int{2, 3, 4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Slice_1(tt.args.s, tt.args.N); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Slice_1() = %v, want %v", got, tt.want)
			}
		})
	}
}
