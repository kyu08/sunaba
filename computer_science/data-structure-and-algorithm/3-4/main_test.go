package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		in     []int
		expect int
	}{
		"0": {
			in:     []int{2, 3, 4, 5},
			expect: 0,
		},
		"1": {
			in:     []int{2, 3, 1, 5},
			expect: 2,
		},
		"2": {
			in:     []int{6, 4, 2, 5, 3},
			expect: 2,
		},
		"3": {
			in:     []int{6, 4, 2, 5, 3},
			expect: 2,
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := minIndex(tt.in)
			if diff := cmp.Diff(tt.expect, got); diff != "" {
				t.Errorf("mismatch. (-expect +got)\n%s", diff)
			}
		})
	}
}
