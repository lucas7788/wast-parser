package lexer

import (
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLexer(t *testing.T) {
	lexer := NewLexer("aaa\"saaa\"()))((")

	for {
		token, err := lexer.Parse()
		if err != nil {
			if err != io.EOF {
				fmt.Printf("read token error: err: %v\n", err)
			}
			return
		}

		fmt.Printf("token: %v\n", token)
	}

}

func skipComment(input string) string {
	lexer := NewLexer(input)
	lexer.SkipComment()
	return string(lexer.buf.Bytes())
}

func TestLineComment(t *testing.T) {
	assert.Equal(t, skipComment(";;"), "")
	assert.Equal(t, skipComment(";;xyz "), "")
	assert.Equal(t, skipComment(";;xaz\nabc"), "abc")
	assert.Equal(t, skipComment(";;xa\nz\nabc"), "z\nabc")
	assert.Equal(t, skipComment(";;x;;a\nz\nabc"), "z\nabc")
}

func TestBlockComment(t *testing.T) {
	assert.Equal(t, skipComment("(;;)"), "")
	assert.Equal(t, skipComment("(; ;)"), "")
	assert.Equal(t, skipComment("(; (;;) ;)"), "")
}

func TestId(t *testing.T) {
	testGetId := func(input string, expected string) {
		lexer := NewLexer(input)
		token, err := lexer.Parse()
		assert.Nil(t, err)
		id, ok := token.(Identifier)
		assert.True(t, ok)

		assert.Equal(t, id.Val, expected)
	}

	testGetId("$x", "$x")
	testGetId("$xyz", "$xyz")
	testGetId("$x_z", "$x_z")
	testGetId("$0^", "$0^")
	testGetId("$0^", "$0^")
	testGetId("$0^ ", "$0^")
}

func TestInteger(t *testing.T) {
	testInteger := func(input string, expected string) {
		lexer := NewLexer(input)
		token, err := lexer.Parse()
		assert.Nil(t, err)
		id, ok := token.(Integer)
		assert.True(t, ok)

		assert.Equal(t, id.Val, expected)
	}

	testInteger("1", "1")
	testInteger("0", "0")
	testInteger("-1", "-1")
	testInteger("+1", "1")
	testInteger("+1_000", "1000")
	testInteger("+1_0_0_0", "1000")
	testInteger("+0x10", "10")
	testInteger("-0x10", "-10")
	testInteger("0x10", "10")

}
