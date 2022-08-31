package slices

import (
	"reflect"
	"testing"
)

func TestSlice_9(t *testing.T) {
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
		{"three elements", args{[]int{1, 2, 3}}, []int{2, 3, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Slice_9(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Slice_9() = %v, want %v", got, tt.want)
			}
		})
	}
}
