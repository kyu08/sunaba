package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		in     []command
		expect []string
	}{
		"insert": {
			in: []command{
				{"insert", "AAA"},
				{"insert", "AAC"},
				{"find", "AAA"},
				{"find", "CCC"},
				{"insert", "CCC"},
				{"find", "CCC"},
			},
			expect: []string{"yes", "no", "yes"},
		},
		"insert-find": {
			in: []command{
				{"insert", "AAC"},
				{"find", "AAC"},
			},
			expect: []string{"yes"},
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := main_(tt.in)

			if diff := cmp.Diff(tt.expect, got); diff != "" {
				t.Errorf("mismatch. (-expect +got)\n%s", diff)
			}
		})
	}
}
