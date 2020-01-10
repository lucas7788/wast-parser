package ast

import "github.com/ontio/wast-parser/parser"

type Wat struct {
	Module Module
}

func (self *Wat) Parse(ps *parser.ParserBuffer) error {
	return ps.Parens(func(ps *parser.ParserBuffer) error {
		return self.Module.Parse(ps)
	})
}
