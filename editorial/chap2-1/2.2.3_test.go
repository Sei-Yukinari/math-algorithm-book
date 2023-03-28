package chap2_1

import (
	"reflect"
	"testing"
)

func TestP2223(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want P2231Result
	}{
		{
			name: "Case 1",
			args: args{
				x: 13,
				y: 14,
			},
			want: P2231Result{
				answer1: 12,
				answer2: 15,
				answer3: 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := P22231(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("P2223() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestP22232(t *testing.T) {
	type args struct {
		a int
		b int
		c int
		d int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case 1",
			args: args{
				a: 8,
				b: 4,
				c: 2,
				d: 2,
			},
			want: 14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := P22232(tt.args.a, tt.args.b, tt.args.c, tt.args.d); got != tt.want {
				t.Errorf("P22232() = %v, want %v", got, tt.want)
			}
		})
	}
}
