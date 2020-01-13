package ast

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/ontio/wast-parser/parser"
)

func TestExpression(t *testing.T) {
	//ps, err := parser.NewParserBuffer(` (i32.const 0) unreachable (nop) )`)
	ps, err := parser.NewParserBuffer(`(if (local.get 0) (then))`)
	assert.Nil(t, err)

	var expr Expression
	err = expr.Parse(ps)
	assert.Nil(t, err)

	fmt.Printf("tokens: %s, err: %v", expr, err)
}
