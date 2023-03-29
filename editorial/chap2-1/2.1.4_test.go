package chap2_2

import (
	"reflect"
	"testing"
)

func TestP214(t *testing.T) {
	type args struct {
		n1 string
		n2 string
	}
	tests := []struct {
		name string
		args args
		want P214Result
	}{
		{
			name: "Case 1",
			args: args{
				n1: "1001",
				n2: "127",
			},
			want: P214Result{
				decimalNumber: 9,
				nNumber: NNumber{
					binaryNumber:  "1111111",
					ternaryNumber: "11201",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := P214(tt.args.n1, tt.args.n2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("P214() = %v, want %v", got, tt.want)
			}
		})
	}
}
