package slices

import (
	"reflect"
	"testing"
)

func TestSlice_14(t *testing.T) {
	type args struct {
		s []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{"empty slice", args{[]int{}}, []int{}},
		{"one element", args{[]int{1}}, []int{1}},
		{"two elements", args{[]int{1, 2}}, []int{2, 1}},
		{"three elements", args{[]int{1, 2, 3}}, []int{2, 1, 3}},
		{"four elements", args{[]int{1, 2, 3, 4}}, []int{2, 1, 4, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Slice_14(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Slice_14() = %v, want %v", got, tt.want)
			}
		})
	}
}
