package main

import "testing"

func TestSqrt(t *testing.T) {
	type args struct {
		x float64
	}

	tests := []struct {
		name string
		args args
		want float64
	}{
		{"sqrt of 4", args{4}, 2},
		{"sqrt of 5", args{5}, 2.23606797749979},
		{"sqrt of 6", args{6}, 2.449489742783178},
		{"sqrt of 7", args{7}, 2.6457513110645907},
		{"sqrt of 8", args{8}, 2.8284271247461903},
		{"sqrt of 9", args{9}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sqrt(tt.args.x); got != tt.want {
				t.Errorf("Sqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRunMain(t *testing.T) {
	main()
}
