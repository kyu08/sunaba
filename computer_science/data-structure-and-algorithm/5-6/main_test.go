package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		k      int
		wn     []int
		expect int
	}{
		"": {
			k:      3,
			wn:     []int{8, 1, 7, 3, 9},
			expect: 10,
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := main_(tt.k, tt.wn)
			if diff := cmp.Diff(tt.expect, got); diff != "" {
				t.Errorf("mismatch. (-expect +got)\n%s", diff)
			}
		})
	}
}
