package slices

import (
	"reflect"
	"testing"
)

func TestSlice_8(t *testing.T) {
	type args struct {
		a []int
		b []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{"empty slice", args{[]int{}, []int{}}, []int{}},
		{"one element", args{[]int{1}, []int{}}, []int{1}},
		{"two elements", args{[]int{1, 2}, []int{}}, []int{1, 2}},
		{"three elements", args{[]int{1, 2, 3}, []int{1}}, []int{2, 3}},
		{"three elements", args{[]int{1, 2, 3}, []int{2}}, []int{1, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Slice_8(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Slice_8() = %v, want %v", got, tt.want)
			}
		})
	}
}
