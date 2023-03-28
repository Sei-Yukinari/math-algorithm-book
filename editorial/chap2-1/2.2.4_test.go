package chap2_1

import "testing"

func TestP2224(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case 1",
			args: args{a: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
			want: 45,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := P2224(tt.args.a); got != tt.want {
				t.Errorf("P2224() = %v, want %v", got, tt.want)
			}
		})
	}
}
