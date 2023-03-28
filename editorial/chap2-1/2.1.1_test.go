package chap2_1

import (
	"reflect"
	"testing"
)

func TestP211(t *testing.T) {
	type args struct {
		input []any
	}
	tests := []struct {
		name string
		args args
		want Result
	}{
		{
			name: "Case 1",
			args: args{
				input: []any{-100, -20, -1.333, 0, 1, "π", 84 / 11, 12.25, 70},
			},
			want: Result{
				integer:         []any{-100, -20, 0, 1, 84 / 11, 70},
				positiveInteger: []any{1, 84 / 11, 70},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := P211(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("P211() = %v, want %v", got, tt.want)
			}
		})
	}
}
