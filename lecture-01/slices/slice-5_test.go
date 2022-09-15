package slices

import (
	"reflect"
	"testing"
)

func TestSlice_5(t *testing.T) {
	type args struct {
		s []int
	}

	tests := []struct {
		name  string
		args  args
		want  []int
		want1 int
	}{
		{"empty slice", args{[]int{}}, []int{}, 0},
		{"one element", args{[]int{1}}, []int{}, 1},
		{"two elements", args{[]int{1, 2}}, []int{2}, 1},
		{"three elements", args{[]int{1, 2, 3}}, []int{2, 3}, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Slice_5(tt.args.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Slice_5() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Slice_5() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
