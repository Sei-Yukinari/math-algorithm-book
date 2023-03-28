package chap2_1

import "testing"

func TestP212(t *testing.T) {
	type args struct {
		a int
		b int
		c int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case 1",
			args: args{
				a: 25,
				b: 4,
				c: 12,
			},
			want: 41,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := P212(tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("P212() = %v, want %v", got, tt.want)
			}
		})
	}
}
