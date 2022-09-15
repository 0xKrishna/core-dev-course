package maps

import (
	"reflect"
	"testing"
)

func TestMap_3(t *testing.T) {
	type args struct {
		s1 []int
		s2 []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{"empty slice", args{[]int{}, []int{}}, []int{}},
		{"one element", args{[]int{1}, []int{1}}, []int{1}},
		{"two elements", args{[]int{1, 2}, []int{1, 2}}, []int{1, 2}},
		{"two elements", args{[]int{1, 2}, []int{2, 3}}, []int{2}},
		{"three elements", args{[]int{1, 2, 3}, []int{2, 3}}, []int{2, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Map_3(tt.args.s1, tt.args.s2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map_3() = %v, want %v", got, tt.want)
			}
		})
	}
}
