package slices

import (
	"reflect"
	"testing"
)

func TestSlice_15(t *testing.T) {
	type args struct {
		s []int
		a []string
	}

	tests := []struct {
		name  string
		args  args
		want  []int
		want1 []int
		want2 []string
	}{
		{"empty slice", args{[]int{}, []string{}}, []int{}, []int{}, []string{}},
		{"one element", args{[]int{1}, []string{"1"}}, []int{1}, []int{1}, []string{"1"}},
		{"two elements", args{[]int{1, 2}, []string{"aab", "aaa"}}, []int{1, 2}, []int{2, 1}, []string{"aaa", "aab"}},
		{"three elements", args{[]int{2, 3, 1}, []string{"b", "c", "a"}}, []int{1, 2, 3}, []int{3, 2, 1}, []string{"a", "b", "c"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := Slice_15(tt.args.s, tt.args.a)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Slice_15() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Slice_15() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("Slice_15() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
