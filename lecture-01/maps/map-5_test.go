package maps

import (
	"reflect"
	"testing"
)

func TestMap_5(t *testing.T) {
	type args struct {
		m map[int]map[int]string
	}

	tests := []struct {
		name string
		args args
		want map[int]map[int]string
	}{
		{"empty map", args{map[int]map[int]string{}}, map[int]map[int]string{}},
		{"one element", args{map[int]map[int]string{1: {1: "a"}}}, map[int]map[int]string{1: {1: "a"}}},
		{"two elements", args{map[int]map[int]string{7: {8: "a", 5: "b"}, 9: {3: "x", 1: "y"}}}, map[int]map[int]string{9: {1: "y", 3: "x"}, 7: {5: "b", 8: "a"}}},
		{"three elements", args{map[int]map[int]string{5: {1: "a"}, 3: {2: "b"}, 9: {3: "c"}}}, map[int]map[int]string{9: {3: "c"}, 5: {1: "a"}, 3: {2: "b"}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Map_5(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map_5() = %v, want %v", got, tt.want)
			}
		})
	}
}
