package slices

import (
	"reflect"
	"testing"
)

func TestSlice_7(t *testing.T) {
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
		{"one element", args{[]int{1}, []int{1}}, []int{1}},
		{"two elements", args{[]int{1, 2}, []int{2, 1}}, []int{1, 2}},
		{"three elements", args{[]int{1, 2, 3}, []int{2, 3, 4}}, []int{1, 2, 3, 4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Slice_7(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Slice_7() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_contains(t *testing.T) {
	type args struct {
		s []int
		e int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{"empty slice", args{[]int{}, 0}, false},
		{"one element", args{[]int{1}, 1}, true},
		{"two elements", args{[]int{1, 2}, 4}, false},
		{"three elements", args{[]int{1, 2, 3}, 3}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := contains(tt.args.s, tt.args.e); got != tt.want {
				t.Errorf("contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
