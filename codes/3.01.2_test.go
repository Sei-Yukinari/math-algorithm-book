package main

import "testing"

func TestC3012(t *testing.T) {
	type args struct {
		n uint
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Case 1",
			args: args{
				n: 11,
			},
			want: true,
		},
		{
			name: "Case 2",
			args: args{
				n: 100,
			},
			want: false,
		},
		{
			name: "Case 3",
			args: args{
				n: 10000000019,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := C3012(tt.args.n); got != tt.want {
				t.Errorf("C3012() = %v, want %v", got, tt.want)
			}
		})
	}
}
