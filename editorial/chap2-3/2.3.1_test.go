package chap2_3

import "testing"

func TestP231(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case 1",
			args: args{
				x: 1,
			},
			want: 1,
		},
		{
			name: "Case 2",
			args: args{
				x: 5,
			},
			want: 125,
		},
		{
			name: "Case 3",
			args: args{
				x: 10,
			},
			want: 1000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := P231(tt.args.x); got != tt.want {
				t.Errorf("P231() = %v, want %v", got, tt.want)
			}
		})
	}
}
