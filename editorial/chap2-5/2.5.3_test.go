package chap2_5

import "testing"

func TestP253(t *testing.T) {
	type args struct {
		n uint
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{
			name: "Case 1",
			args: args{
				n: 20,
			},
			want: 2432902008176640000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := P253(tt.args.n); got != tt.want {
				t.Errorf("P253() = %v, want %v", got, tt.want)
			}
		})
	}
}
