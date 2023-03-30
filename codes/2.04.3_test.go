package main

import "testing"

func TestC2043(t *testing.T) {
	type args struct {
		n int
		s int
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
				s: 4,
			},
			want: 6,
		},
		{
			name: "Case 2",
			args: args{
				n: 1000,
				s: 100,
			},
			want: 4950,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := C2043(tt.args.n, tt.args.s); got != tt.want {
				t.Errorf("C2043() = %v, want %v", got, tt.want)
			}
		})
	}
}
