package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_main(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		q      int
		input  []task
		expect []task
	}{
		"解答": {
			q:      100,
			input:  []task{{"p1", 150}, {"p2", 80}, {"p3", 200}, {"p4", 350}, {"p5", 20}},
			expect: []task{{"p2", 180}, {"p5", 400}, {"p1", 450}, {"p3", 550}, {"p4", 800}},
		},
		"短めパターン": {
			q:      100,
			input:  []task{{"p1", 10}, {"p2", 80}, {"p3", 200}},
			expect: []task{{"p1", 10}, {"p2", 90}, {"p3", 290}},
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := main_(tt.q, tt.input)
			if diff := cmp.Diff(tt.expect, got); diff != "" {
				t.Errorf("mismatch. (-expect +got)\n%s", diff)
			}
		})
	}
}
