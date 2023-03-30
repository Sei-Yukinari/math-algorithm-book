package chap3_1

import (
	"reflect"
	"testing"
)

func TestP312(t *testing.T) {
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
				n: 286,
			},
			want: []int{2, 11, 13},
		},
		{
			name: "Case 2",
			args: args{
				n: 20211225,
			},
			want: []int{3, 5, 5, 31, 8693},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := P312(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("P312() = %v, want %v", got, tt.want)
			}
		})
	}
}
