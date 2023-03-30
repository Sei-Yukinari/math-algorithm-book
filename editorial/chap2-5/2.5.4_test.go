package chap2_5

import (
	"reflect"
	"testing"
)

func TestP254(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Case 1",
			args: args{
				n: 20,
			},
			want: []int{2, 3, 5, 7, 11, 13, 17, 19},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := P254(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("P254() = %v, want %v", got, tt.want)
			}
		})
	}
}
