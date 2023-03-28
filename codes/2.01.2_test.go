package main

import "testing"

func TestC2012(t *testing.T) {
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
				a1: 10,
				a2: 20,
				a3: 50,
			},
			want: 80,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := C2012(tt.args.a1, tt.args.a2, tt.args.a3); got != tt.want {
				t.Errorf("C2012() = %v, want %v", got, tt.want)
			}
		})
	}
}
