package main

import "testing"

func TestC2041(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case 1",
			args: args{
				n: 3,
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := C2041(tt.args.n); got != tt.want {
				t.Errorf("C2041() = %v, want %v", got, tt.want)
			}
		})
	}
}
