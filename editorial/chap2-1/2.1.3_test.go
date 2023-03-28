package chap2_1

import "testing"

func TestP213(t *testing.T) {
	type args struct {
		a1 int
		a2 int
		a3 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case 1",
			args: args{
				a1: 2,
				a2: 8,
				a3: 8,
			},
			want: 128,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := P213(tt.args.a1, tt.args.a2, tt.args.a3); got != tt.want {
				t.Errorf("P213() = %v, want %v", got, tt.want)
			}
		})
	}
}
