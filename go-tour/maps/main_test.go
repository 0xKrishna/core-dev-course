package main

import (
	"reflect"
	"testing"
)

func TestWordCount(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{"empty string", args{""}, map[string]int{}},
		{"one word", args{"word"}, map[string]int{"word": 1}},
		{"two words", args{"word one"}, map[string]int{"word": 1, "one": 1}},
		{"three words", args{"word one two"}, map[string]int{"word": 1, "one": 1, "two": 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WordCount(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WordCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRunMain(t *testing.T) {
	main()
}
