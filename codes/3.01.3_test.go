package main

import (
	"reflect"
	"testing"
)

func TestC3013(t *testing.T) {
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
				n: 100,
			},
			want: []int{1, 2, 4, 5, 10, 20, 25, 50, 100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := C3013(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("C3013() = %v, want %v", got, tt.want)
			}
		})
	}
}
