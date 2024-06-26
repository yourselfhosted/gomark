package parser

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/usememos/gomark/ast"
	"github.com/usememos/gomark/parser/tokenizer"
	"github.com/usememos/gomark/restore"
)

func TestBoldItalicParser(t *testing.T) {
	tests := []struct {
		text       string
		boldItalic ast.Node
	}{
		{
			text:       "*Hello world!",
			boldItalic: nil,
		},
		{
			text:       "*** Hello * *",
			boldItalic: nil,
		},
		{
			text:       "*** Hello **",
			boldItalic: nil,
		},
		{
			text: "***Hello***",
			boldItalic: &ast.BoldItalic{
				Symbol:  "*",
				Content: "Hello",
			},
		},
		{
			text: "*** Hello ***",
			boldItalic: &ast.BoldItalic{
				Symbol:  "*",
				Content: " Hello ",
			},
		},
	}

	for _, test := range tests {
		tokens := tokenizer.Tokenize(test.text)
		node, _ := NewBoldItalicParser().Match(tokens)
		require.Equal(t, restore.Restore([]ast.Node{test.boldItalic}), restore.Restore([]ast.Node{node}))
	}
}
