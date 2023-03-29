package chap2_3

import "testing"

func TestP2331(t *testing.T) {
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
			args: args{x: 3},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := P2331(tt.args.x); got != tt.want {
				t.Errorf("P2331() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestP2332(t *testing.T) {
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
			args: args{x: 3},
			want: 1000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := P2332(tt.args.x); got != tt.want {
				t.Errorf("P2332() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestP2333(t *testing.T) {
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
			args: args{x: 16},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := P2333(tt.args.x); got != tt.want {
				t.Errorf("P2333() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestP2334(t *testing.T) {
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
			args: args{x: 16},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := P2334(tt.args.x); got != tt.want {
				t.Errorf("P2334() = %v, want %v", got, tt.want)
			}
		})
	}
}
