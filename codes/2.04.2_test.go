package main

import "testing"

func TestC2042(t *testing.T) {
	type args struct {
		n int
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case 1",
			args: args{
				n: 15,
				x: 3,
				y: 5,
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := C2042(tt.args.n, tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("C2042() = %v, want %v", got, tt.want)
			}
		})
	}
}
