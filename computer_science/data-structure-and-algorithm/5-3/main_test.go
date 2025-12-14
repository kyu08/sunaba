package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMain_(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		s      []int
		t      []int
		expect int
	}{
		"": {
			s:      []int{1, 2, 3, 4, 5},
			t:      []int{3, 4, 1},
			expect: 3,
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := main_(tt.s, tt.t)
			if diff := cmp.Diff(tt.expect, got); diff != "" {
				t.Errorf("mismatch. (-expect +got)\n%s", diff)
			}
		})
	}
}

func TestBinarySearch(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		s      []int
		num    int
		expect bool
	}{
		"true(mid)": {
			s:      []int{1, 2, 3, 4, 5},
			num:    3,
			expect: true,
		},
		"true(mid以下)": {
			s:      []int{1, 2, 3, 4, 5},
			num:    2,
			expect: true,
		},
		"true(mid以上)": {
			s:      []int{1, 2, 3, 4, 5},
			num:    4,
			expect: true,
		},
		"false(min以下)": {
			s:      []int{1, 2, 3, 4, 5},
			num:    0,
			expect: false,
		},
		"false(max以上)": {
			s:      []int{1, 2, 3, 4, 5},
			num:    100,
			expect: false,
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := binarySearch(tt.s, tt.num)
			if diff := cmp.Diff(tt.expect, got); diff != "" {
				t.Errorf("mismatch. (-expect +got)\n%s", diff)
			}
		})
	}
}
