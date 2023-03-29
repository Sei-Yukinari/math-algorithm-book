package chap2_3

import "testing"

func TestP2341(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case 1",
			args: args{x: 20},
			want: 40,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := P2341(tt.args.x); got != tt.want {
				t.Errorf("P2341() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestP2342(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Case 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			P2342()
		})
	}
}
