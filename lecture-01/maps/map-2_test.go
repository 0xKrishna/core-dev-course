package maps

import (
	"reflect"
	"testing"
)

func TestMap_2(t *testing.T) {
	type args struct {
		s []int
	}

	tests := []struct {
		name string
		args args
		want map[int]struct{}
	}{
		{"empty slice", args{[]int{}}, map[int]struct{}{}},
		{"one element", args{[]int{1}}, map[int]struct{}{1: {}}},
		{"two elements", args{[]int{1, 2}}, map[int]struct{}{1: {}, 2: {}}},
		{"three elements", args{[]int{1, 2, 3}}, map[int]struct{}{1: {}, 2: {}, 3: {}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Map_2(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map_2() = %v, want %v", got, tt.want)
			}
		})
	}
}
