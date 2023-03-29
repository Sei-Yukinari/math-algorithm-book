package chap2_3

import (
	"reflect"
	"testing"
)

func TestP2321(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "Case 1",
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := P2321(); got != tt.want {
				t.Errorf("P2321() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestP2322(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "Case 1",
			want: 3162,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := P2322(); got != tt.want {
				t.Errorf("P2322() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestP2323(t *testing.T) {
	tests := []struct {
		name string
		want []int
	}{
		{
			name: "Case 1",
			want: []int{20, 21},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := P2323(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("P2323() = %v, want %v", got, tt.want)
			}
		})
	}
}
