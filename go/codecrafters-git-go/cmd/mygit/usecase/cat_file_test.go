package usecase

import (
	"errors"
	"github.com/codecrafters-io/git-starter-go/cmd/mygit/testable"
	"testing"
)

func TestCatFileParamValidate(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		Hash string
		Want error
	}{
		"40文字のときnilを返す": {
			Hash: "1234567890123456789012345678901234567890",
			Want: nil,
		},
		"39文字のときエラーを返す": {
			Hash: "123456789012345678901234567890123456789",
			Want: errors.New("invalid hash format."),
		},
	}
	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			p := CatFileParam{
				Hash: &tt.Hash,
			}
			got := p.validate()
			want := tt.Want
			if diff := testable.CompareError(t, want, got); diff != "" {
				t.Errorf("mismatch. (-expect +got)\n%s", diff)
			}
		})
	}
}
