package main

import "testing"

func TestC2013(t *testing.T) {
	type args struct {
		n []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case 1",
			args: args{
				[]int{3, 1, 4, 1, 5},
			},
			want: 14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := C2013(tt.args.n); got != tt.want {
				t.Errorf("C2013() = %v, want %v", got, tt.want)
			}
		})
	}
}
