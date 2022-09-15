package slices

import (
	"reflect"
	"testing"
)

func TestSlice_10(t *testing.T) {
	type args struct {
		s []int
		i int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{"empty slice", args{[]int{}, 0}, []int{}},
		{"one element", args{[]int{1}, 0}, []int{1}},
		{"two elements", args{[]int{1, 2}, 1}, []int{2, 1}},
		{"three elements", args{[]int{1, 2, 3}, 2}, []int{3, 1, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Slice_10(tt.args.s, tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Slice_10() = %v, want %v", got, tt.want)
			}
		})
	}
}
