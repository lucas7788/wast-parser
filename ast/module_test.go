package ast

import (
	"fmt"
	"testing"

	"github.com/ontio/wast-parser/parser"
)

func TestImportGlobal_ImportType(t *testing.T) {

	ps, err := parser.NewParserBuffer(`
  import "test" "table-10-inf" (global i32)))
`)

	if err != nil {
		fmt.Printf("error: %s", err)
	}

	fmt.Printf("tokens: %v", ps)
}
