package main

import "testing"

func Test_C2011(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case 1",
			args: args{
				num: 2,
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := C2011(tt.args.num); got != tt.want {
				t.Errorf("orangeAndApple() = %v, want %v", got, tt.want)
			}
		})
	}
}
