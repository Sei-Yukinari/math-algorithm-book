package chap2_2

import "testing"

func TestP221(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case 10,000",
			args: args{
				index: 4,
			},
			want: 10000,
		},
		{
			name: "Case 100,000,000",
			args: args{
				index: 8,
			},
			want: 100000000,
		},
		{
			name: "Case 10,000",
			args: args{
				index: 12,
			},
			want: 1000000000000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := P221(tt.args.index); got != tt.want {
				t.Errorf("P221() = %v, want %v", got, tt.want)
			}
		})
	}
}
