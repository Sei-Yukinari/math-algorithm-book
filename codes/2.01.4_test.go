package main

import "testing"

func TestC2014(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Case 1",
			args: args{
				n: 3,
			},
			want: "11",
		},
		{
			name: "Case 2",
			args: args{
				n: 5,
			},
			want: "101",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := C2014(tt.args.n); got != tt.want {
				t.Errorf("C2014() = %v, want %v", got, tt.want)
			}
		})
	}
}
