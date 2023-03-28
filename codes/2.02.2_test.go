package main

import (
	"reflect"
	"testing"
)

func TestC2022(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want C2022Result
	}{
		{
			name: "Case 1",
			args: args{
				a: 11,
				b: 14,
			},
			want: C2022Result{
				and: 10,
				or:  15,
				xor: 5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := C2022(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("C2022() = %v, want %v", got, tt.want)
			}
		})
	}
}
