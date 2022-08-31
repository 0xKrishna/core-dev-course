package maps

import (
	"reflect"
	"testing"
)

func TestMap_1(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{"empty string", args{""}, map[string]int{}},
		{"one character", args{"a"}, map[string]int{"a": 1}},
		{"two characters", args{"ab"}, map[string]int{"a": 1, "b": 1}},
		{"three characters", args{"abc"}, map[string]int{"a": 1, "b": 1, "c": 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Map_1(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map_1() = %v, want %v", got, tt.want)
			}
		})
	}
}
