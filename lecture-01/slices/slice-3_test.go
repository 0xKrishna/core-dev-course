package slices

import (
	"reflect"
	"testing"
)

func TestSlice_3(t *testing.T) {
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
		{"two elements", args{[]int{1, 2}, 3}, []int{4, 2}},
		{"three elements", args{[]int{1, 2, 3}, 4}, []int{5, 2, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Slice_3(tt.args.s, tt.args.N); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Slice_3() = %v, want %v", got, tt.want)
			}
		})
	}
}
