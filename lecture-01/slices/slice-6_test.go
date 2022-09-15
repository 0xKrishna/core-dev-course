package slices

import (
	"reflect"
	"testing"
)

func TestSlice_6(t *testing.T) {
	type args struct {
		s []int
		i int
	}

	tests := []struct {
		name  string
		args  args
		want  []int
		want1 int
	}{
		{"empty slice", args{[]int{}, 0}, []int{}, 0},
		{"one element", args{[]int{1}, 0}, []int{}, 1},
		{"two elements", args{[]int{1, 2}, 1}, []int{1}, 2},
		{"three elements", args{[]int{1, 2, 3}, 1}, []int{1, 3}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Slice_6(tt.args.s, tt.args.i)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Slice_6() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Slice_6() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
