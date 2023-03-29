package chap2_2

import (
	"reflect"
	"testing"
)

func TestP222_1(t *testing.T) {
	type args struct {
		n1 int
		n2 int
	}
	tests := []struct {
		name string
		args args
		want P2221Result
	}{
		{
			name: "Case 1",
			args: args{
				n1: 841,
				n2: 29,
			},
			want: P2221Result{
				answer1: 29,
				answer2: 841,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := P2221(tt.args.n1, tt.args.n2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("P222_1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestP2222(t *testing.T) {
	type args struct {
		x1 float64
		y1 float64
		x2 int
		y2 int
	}
	tests := []struct {
		name string
		args args
		want P2222Result
	}{
		{
			name: "Case 1",
			args: args{
				x1: 5.0,
				y1: 1024.0,
				x2: 4,
				y2: 5,
			},
			want: P2222Result{
				answer1: 4,
				answer2: 1024,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := P2222(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("P2222() = %v, want %v", got, tt.want)
			}
		})
	}
}
