package chap2_3

import "testing"

func TestP2351(t *testing.T) {
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
			args: args{x: 1000000},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := P2351(tt.args.x); got != tt.want {
				t.Errorf("P2351() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestP2352(t *testing.T) {
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
			args: args{n: 100},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := P2352(tt.args.n); got != tt.want {
				t.Errorf("P2352() = %v, want %v", got, tt.want)
			}
		})
	}
}
