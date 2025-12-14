package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		in     []int
		expect [][]int
	}{
		"": {
			in: []int{5, 3, 4, 2},
			expect: [][]int{
				{5, 3, 4, 2},
				{3, 5, 4, 2},
				{3, 4, 5, 2},
				{3, 4, 2, 5},
				{3, 2, 4, 5},
				{2, 3, 4, 5},
			},
		},
		// "": {
		// 	in: []int{5, 2, 4, 6, 1, 3},
		// 	expect: [][]int{
		// 		{5, 2, 4, 6, 1, 3},
		// 		{2, 5, 4, 6, 1, 3},
		// 		{2, 4, 5, 6, 1, 3},
		// 		{2, 4, 5, 6, 1, 3},
		// 		{1, 2, 4, 5, 6, 3},
		// 		{1, 2, 3, 4, 5, 6},
		// 	},
		// },
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := insertionSort(tt.in, len(tt.in))
			if diff := cmp.Diff(tt.expect, got); diff != "" {
				t.Errorf("mismatch. (-expect +got)\n%s", diff)
			}
		})
	}
}
