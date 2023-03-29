package main

import (
	"reflect"
	"testing"
)

func TestC2031(t *testing.T) {
	tests := []struct {
		name string
		want C2031Result
	}{
		{
			name: "Case 1",
			want: C2031Result{
				a1: 2021,
				a2: 1501,
				a3: 1502,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := C2031(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("C2031() = %v, want %v", got, tt.want)
			}
		})
	}
}
