package parser

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/usememos/gomark/ast"
	"github.com/usememos/gomark/parser/tokenizer"
	"github.com/usememos/gomark/restore"
)

func TestSpoilerParser(t *testing.T) {
	tests := []struct {
		text    string
		spoiler ast.Node
	}{
		{
			text:    "*Hello world!",
			spoiler: nil,
		},
		{
			text: "||Hello||",
			spoiler: &ast.Spoiler{
				Content: "Hello",
			},
		},
	}

	for _, test := range tests {
		tokens := tokenizer.Tokenize(test.text)
		node, _ := NewSpoilerParser().Match(tokens)
		require.Equal(t, restore.Restore([]ast.Node{test.spoiler}), restore.Restore([]ast.Node{node}))
	}
}
